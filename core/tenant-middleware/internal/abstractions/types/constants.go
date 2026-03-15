package types

const (
	SearchTenantsQuery = `SELECT id, name, subdomain, db_host, db_port, db_name, db_user, db_password, db_ssl_mode, settings, status_id, created_at FROM tenants WHERE subdomain = $1 LIMIT 1`
	SearchTenantById   = `SELECT id, name, subdomain, db_host, db_port, db_name, db_user, db_password, db_ssl_mode, settings, status_id FROM tenants WHERE id = $1 LIMIT 1`
	UpdateTenantStatus = `UPDATE tenants SET status_id = $1 WHERE id = $2`
)
