package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types/enums"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/business"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/config"
	asigna_multitenancy "github.com/samforedev/asignads/lib/asigna-multitenancy"
)

type Server struct {
	engine   *gin.Engine
	cfg      *config.Config
	resolver *business.TenantResolver
	manager  *business.TenantManager
}

func NewServer(cfg *config.Config, resolver *business.TenantResolver, manager *business.TenantManager) *Server {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &Server{
		engine:   gin.Default(),
		cfg:      cfg,
		resolver: resolver,
		manager:  manager,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	admin := s.engine.Group("/admin/tenants")
	{
		admin.PATCH("/:id/status", func(c *gin.Context) {
			id := c.Param("id")
			var input struct {
				// Gin usará el UnmarshalJSON que acabamos de crear
				Status enums.TenantStatus `json:"status" binding:"required"`
			}

			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
				return
			}

			// El Manager actualiza DB y limpia Redis
			err := s.manager.ChangeStatus(c.Request.Context(), id, input.Status)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Tenant status updated and cache invalidated successfully",
			})
		})
	}

	s.engine.NoRoute(business.TenantLoader(s.resolver), func(c *gin.Context) {
		tenantID, _ := c.Get(string(asigna_multitenancy.TenantIDKey))

		if tenantID == nil || tenantID == "" {
			return
		}

		target := s.cfg.BackendTargetURL
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid proxy target"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Request.URL.Path // Mantiene "/users"
			req.Host = remote.Host

			req.Header.Set("X-Tenant-ID", tenantID.(string))
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.cfg.Port)
}
