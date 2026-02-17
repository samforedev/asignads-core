package types

const (
	SearchTenantsQuery = `SELECT id, name, subdomain, db_dsn, status_id, created_at FROM tenants WHERE subdomain = $1 LIMIT 1`
	SearchTenantById   = `SELECT id, name, subdomain, db_dsn, status_id FROM tenants WHERE id = $1 LIMIT 1`
	UpdateTenantStatus = `UPDATE tenants SET status_id = $1 WHERE id = $2`
)
