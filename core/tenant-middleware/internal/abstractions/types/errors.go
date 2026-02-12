package types

import "errors"

var (
	ErrTenantNotFound = errors.New("tenant not found")
	ErrTenantInactive = errors.New("tenant is inactive")
)
