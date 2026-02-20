package asigna_multitenancy

import (
	"context"

	baseentitiesconstants "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/constant"
)

func ExtractTenantID(ctx context.Context) string {
	if id, ok := ctx.Value(baseentitiesconstants.TenantIDKey).(string); ok {
		return id
	}
	return ""
}
