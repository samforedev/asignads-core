package domain

import (
	"time"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types/enums"
)

type Tenant struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Subdomain string             `json:"subdomain"`
	DBDSN     string             `json:"db_dsn"`
	Status    enums.TenantStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
}

func (t *Tenant) IsActive() bool {
	return t.Status == enums.Active
}
