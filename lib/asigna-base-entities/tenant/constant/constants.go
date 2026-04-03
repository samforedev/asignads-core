package constant

type ContextKey string

const TenantIDKey ContextKey = "x-tenant-id"

const (
	SearchTenantBySubdomain = `SELECT t.id, t.name as tenant_name, t.subdomain, t.status_id, db.host, db.port, db.db_name, db.db_user, db.db_password, db.ssl_mode from tenants t inner join tenant_database_info db on t.id = db.tenant_id WHERE t.subdomain = $1 LIMIT 1`
)
