package business

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	baseentities "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/constant"
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

		ctx := context.WithValue(c.Request.Context(), baseentities.TenantIDKey, tenant.ID)
		c.Request = c.Request.WithContext(ctx)
		c.Set(string(baseentities.TenantIDKey), tenant.ID)
		c.Writer.Header().Set(string(baseentities.TenantIDKey), tenant.ID)
		c.Next()
	}
}
