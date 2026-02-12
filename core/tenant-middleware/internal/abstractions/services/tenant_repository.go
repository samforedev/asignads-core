package services

import (
	"context"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/domain"
)

type TenantRepository interface {
	GetBySubDomain(ctx context.Context, subdomain string) (*domain.Tenant, error)
	SaveInCache(ctx context.Context, tenant *domain.Tenant) error
}
