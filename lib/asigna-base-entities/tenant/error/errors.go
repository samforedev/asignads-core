package error

import "errors"

var (
	ErrTenantNotFound       = errors.New("tenant not found")
	ErrTenantInactive       = errors.New("tenant is inactive")
	ErrMissingTenantContext = errors.New("tenant information (ID/DSN) missing in context")
	ErrConnectionFailed     = errors.New("failed to establish connection to tenant database")
)
