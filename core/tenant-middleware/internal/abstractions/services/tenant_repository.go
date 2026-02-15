package services

import (
	"context"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/domain"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types/enums"
)

type TenantRepository interface {
	GetById(ctx context.Context, id string) (*domain.Tenant, error)
	UpdateStatus(ctx context.Context, id string, status enums.TenantStatus) error
	GetBySubDomain(ctx context.Context, subdomain string) (*domain.Tenant, error)
	SaveInCache(ctx context.Context, tenant *domain.Tenant) error
}
