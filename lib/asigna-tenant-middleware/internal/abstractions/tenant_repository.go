package abstractions

import (
	"context"

	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
)

type TenantRepository interface {
	GetBySubdomain(ctx context.Context, subdomain string) (*domain.Tenant, error)
	SaveInCache(ctx context.Context, tenant *domain.Tenant) error
}
