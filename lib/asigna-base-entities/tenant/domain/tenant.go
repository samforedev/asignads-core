package domain

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        uuid.UUID
	Name      string
	Subdomain string
	StatusId  TenantStatus
	CreatedAt time.Time
	UpdatedAt time.Time

	Contact  ContactInfo
	Database DatabaseConfig
}

func (tenant *Tenant) IsActive() bool {
	return tenant.StatusId == StatusActive
}
