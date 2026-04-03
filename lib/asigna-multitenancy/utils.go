package asigna_multitenancy

import "github.com/google/uuid"

func GetTenantDSNKey(tenantID uuid.UUID) string {
	return "tenant:" + tenantID.String() + ":dsn"
}
