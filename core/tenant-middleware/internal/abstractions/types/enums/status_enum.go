package enums

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TenantStatus int

const (
	Active       TenantStatus = 1
	Suspended    TenantStatus = 2
	Provisioning TenantStatus = 3
	Deleted      TenantStatus = 4
)

func (s TenantStatus) String() string {
	names := map[TenantStatus]string{
		Active:       "ACTIVE",
		Suspended:    "SUSPENDED",
		Provisioning: "PROVISIONING",
		Deleted:      "DELETED",
	}

	if name, ok := names[s]; ok {
		return name
	}
	return "UNKNOWN"
}

func (s TenantStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TenantStatus) UnmarshalJSON(b []byte) error {
	var statusStr string
	if err := json.Unmarshal(b, &statusStr); err != nil {
		return err
	}
	switch strings.ToUpper(statusStr) {
	case "ACTIVE":
		*s = Active
	case "SUSPENDED":
		*s = Suspended
	case "PROVISIONING":
		*s = Provisioning
	case "DELETED":
		*s = Deleted
	default:
		return fmt.Errorf("invalid tenant status: %s", statusStr)
	}
	return nil
}
