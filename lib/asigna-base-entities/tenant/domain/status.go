package domain

type TenantStatus int

const (
	StatusActive       TenantStatus = 1
	StatusSuspended    TenantStatus = 2
	StatusProvisioning TenantStatus = 3
	StatusDeleted      TenantStatus = 4
)

func (s TenantStatus) String() string {
	switch s {
	case StatusActive:
		return "ACTIVE"
	case StatusSuspended:
		return "SUSPENDED"
	case StatusProvisioning:
		return "PROVISIONING"
	case StatusDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}
