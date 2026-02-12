package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/business"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/config"
	asigna_multitenancy "github.com/samforedev/asignads/lib/asigna-multitenancy"
)

type Server struct {
	engine   *gin.Engine
	cfg      *config.Config
	resolver *business.TenantResolver
}

func NewServer(cfg *config.Config, resolver *business.TenantResolver) *Server {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &Server{
		engine:   gin.Default(),
		cfg:      cfg,
		resolver: resolver,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	s.engine.NoRoute(business.TenantLoader(s.resolver), func(c *gin.Context) {
		tenantID, _ := c.Get(string(asigna_multitenancy.TenantIDKey))

		if tenantID == nil || tenantID == "" {
			return
		}

		target := "http://127.0.0.1:8081"
		remote, _ := url.Parse(target)

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
