package enum

type TenantStatus int

const (
	ACTIVE      TenantStatus = 1
	SUSPENDED   TenantStatus = 2
	PROVISIONED TenantStatus = 3
	INACTIVE    TenantStatus = 4
	DELETED     TenantStatus = 5
)

func (e TenantStatus) String() string {
	names := map[TenantStatus]string{
		ACTIVE:      "ACTIVE",
		SUSPENDED:   "SUSPENDED",
		PROVISIONED: "PROVISIONED",
		INACTIVE:    "INACTIVE",
		DELETED:     "DELETED",
	}

	if name, ok := names[e]; ok {
		return name
	}
	return "UNKNOWN"
}
