package domain

import (
	"time"

	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/enum"
)

type Tenant struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Subdomain string            `json:"subdomain"`
	DBDSN     string            `json:"db_dsn"`
	Status    enum.TenantStatus `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
}

func (t *Tenant) IsActive() bool {
	return t.Status == enum.ACTIVE
}
