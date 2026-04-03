package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	baseentitiesconstants "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/constant"
	"github.com/samforedev/asignads/lib/asigna-tenant-middleware/internal/business"
	"github.com/samforedev/asignads/lib/asigna-tenant-middleware/internal/config"
)

type Server struct {
	engine   *gin.Engine
	cfg      *config.Config
	resolver *business.TenantResolver
	manager  *business.TenantManager
}

func (s *Server) setUpRoutes() {
	s.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	s.engine.NoRoute(business.TenantLoader(s.resolver), func(c *gin.Context) {
		tenantID, _ := c.Get(string(baseentitiesconstants.TenantIDKey))

		if tenantID == nil || tenantID == "" {
			return
		}

		target := ""
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid proxy target"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Request.URL.Path
			req.Host = remote.Host

			req.Header.Set("X-Tenant-ID", tenantID.(string))
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.cfg.AppPort)
}

func NewServer(cfg *config.Config, resolver *business.TenantResolver, manager *business.TenantManager) *Server {
	if cfg.AppEnvironment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := &Server{
		engine:   gin.Default(),
		cfg:      cfg,
		resolver: resolver,
		manager:  manager,
	}
	server.setUpRoutes()
	return server
}
