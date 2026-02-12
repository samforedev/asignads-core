package asigna_multitenancy

import "context"

func ExtractTenantID(ctx context.Context) string {
	if id, ok := ctx.Value(TenantIDKey).(string); ok {
		return id
	}
	return ""
}
