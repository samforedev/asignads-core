package types

type ContextKey string

const SearchTenantsQuery = `SELECT id, name, subdomain, db_dsn, status, created_at FROM tenants WHERE subdomain = $1 LIMIT 1`

const (
	TenantIDKey  ContextKey = "x-tenant-id"
	TenantDSNKey ContextKey = "x-tenant-db-dsn"
)
