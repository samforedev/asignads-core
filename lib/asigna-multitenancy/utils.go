package asigna_multitenancy

func GetTenantDSNKey(tenantID string) string {
	return "tenant:" + tenantID + ":dsn"
}
