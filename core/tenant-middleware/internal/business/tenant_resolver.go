package business

import (
	"context"
	"fmt"
	"strings"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/domain"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/services"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types"
)

type TenantResolver struct {
	cache services.TenantRepository
	db    services.TenantRepository
}

func NewTenantResolver(cache services.TenantRepository, db services.TenantRepository) *TenantResolver {
	return &TenantResolver{
		cache: cache,
		db:    db,
	}
}

func (r *TenantResolver) Resolve(ctx context.Context, host string) (*domain.Tenant, error) {
	subdomain := r.extractSubDomain(host)

	tenant, err := r.cache.GetBySubDomain(ctx, subdomain)
	if err == nil {
		if !tenant.IsActive() {
			return nil, types.ErrTenantInactive
		}
		return tenant, nil
	}

	tenant, err = r.db.GetBySubDomain(ctx, subdomain)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", types.ErrTenantNotFound, subdomain)
	}

	if !tenant.IsActive() {
		return nil, types.ErrTenantInactive
	}

	go func(t *domain.Tenant) {
		_ = r.cache.SaveInCache(context.Background(), t)
	}(tenant)

	return tenant, nil
}

func (r *TenantResolver) extractSubDomain(host string) string {
	hostParts := strings.Split(host, ".")[0]
	return strings.ToLower(hostParts)
}
