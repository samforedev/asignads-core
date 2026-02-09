package business

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types"
)

func TenantLoader(resolver *TenantResolver) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		tenant, err := resolver.Resolve(c.Request.Context(), host)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Tenant not found or inactive",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), types.TenantIDKey, tenant.ID)
		ctx = context.WithValue(ctx, types.TenantDSNKey, tenant.DBDSN)
		c.Request = c.Request.WithContext(ctx)

		c.Set(string(types.TenantIDKey), tenant.ID)
		c.Set(string(types.TenantDSNKey), tenant.DBDSN)

		c.Writer.Header().Set(string(types.TenantIDKey), tenant.ID)
		c.Next()
	}
}
