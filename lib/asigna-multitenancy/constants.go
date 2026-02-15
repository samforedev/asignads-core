package asigna_multitenancy

import "errors"

type ContextKey string

const (
	TenantIDKey  ContextKey = "x-tenant-id"
	TenantDSNKey ContextKey = "x-tenant-db-dsn"
)

var (
	ErrMissingTenantContext = errors.New("tenant information (ID/DSN) missing in context")
	ErrConnectionFailed     = errors.New("failed to establish connection to tenant database")
)
