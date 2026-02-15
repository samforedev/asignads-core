package asigna_multitenancy

import (
	"context"

	"github.com/gin-gonic/gin"
)

func HttpToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Tenant-ID")

		// Aquí es donde "mapeamos" de nuevo a las constantes internas del SDK
		ctx := context.WithValue(c.Request.Context(), TenantIDKey, id)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
