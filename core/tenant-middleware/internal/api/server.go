package api

import (
	"github.com/gin-gonic/gin"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/business"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/config"
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
		tenantID, _ := c.Get(string(types.TenantIDKey))
		c.JSON(200, gin.H{
			"message":   "Tenant identified via fallback",
			"tenant_id": tenantID,
			"path":      c.Request.URL.Path,
		})
	})
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.cfg.Port)
}
