package enums

type TenantStatus int

const (
	TenantStatusActive TenantStatus = iota
	TenantStatusInactive
)

func (t TenantStatus) String() string {
	switch t {
	case TenantStatusActive:
		return "ACTIVE"
	case TenantStatusInactive:
		return "INACTIVE"
	default:
		return "ACTIVE"
	}
}
