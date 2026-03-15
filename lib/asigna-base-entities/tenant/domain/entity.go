package domain

import (
	"time"

	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/enum"
)

type Tenant struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Subdomain    string                 `json:"subdomain"`
	NIT          string                 `json:"nit"`
	ContactName  string                 `json:"contact_name"`
	ContactEmail string                 `json:"contact_email"`
	ContactPhone string                 `json:"contact_phone"`
	Address      string                 `json:"address"`
	City         string                 `json:"city"`
	Department   string                 `json:"department"`
	DBHost       string                 `json:"db_host"`
	DBPort       int                    `json:"db_port"`
	DBName       string                 `json:"db_name"`
	DBUser       string                 `json:"db_user"`
	DBPassword   string                 `json:"-"`
	DBSSLMode    string                 `json:"db_ssl_mode"`
	Settings     map[string]interface{} `json:"settings"`
	Status       enum.TenantStatus      `json:"status"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}

func (t *Tenant) IsActive() bool {
	return t.Status == enum.ACTIVE
}
