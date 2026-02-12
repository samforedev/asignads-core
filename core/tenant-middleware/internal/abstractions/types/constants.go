package types

const SearchTenantsQuery = `SELECT id, name, subdomain, db_dsn, status, created_at FROM tenants WHERE subdomain = $1 LIMIT 1`
