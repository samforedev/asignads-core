package enums

import (
	"encoding/json"
	"strings"
)

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

func (t *TenantStatus) UnmarshalJSON(b []byte) error {
	// Eliminamos las comillas del string JSON
	s := strings.Trim(string(b), "\"")

	switch strings.ToUpper(s) {
	case "ACTIVE":
		*t = TenantStatusActive
	case "INACTIVE":
		*t = TenantStatusInactive
	default:
		// Opcional: manejar error si el status no existe
		*t = TenantStatusActive
	}
	return nil
}

func (t TenantStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}
