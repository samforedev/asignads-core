package business

import (
	"context"
	"fmt"
	"strings"

	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	baseentitieserr "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/error"
	"github.com/samforedev/asignads/lib/asigna-tenant-middleware/internal/abstractions"
)

type TenantResolver struct {
	cache abstractions.TenantRepository
	db    abstractions.TenantRepository
}

func (r *TenantResolver) extractSubDomain(host string) string {
	hostParts := strings.Split(host, ".")[0]
	return strings.ToLower(hostParts)
}

func (r *TenantResolver) validateAndAsyncCache(t *domain.Tenant, shouldCache bool) error {
	if !t.IsActive() {
		return baseentitieserr.ErrTenantInactive
	}

	if shouldCache {
		go func(t *domain.Tenant) {
			_ = r.cache.SaveInCache(context.Background(), t)
		}(t)
	}
	return nil
}

func (r *TenantResolver) Resolve(ctx context.Context, host string) (*domain.Tenant, error) {
	subdomain := r.extractSubDomain(host)
	tenant, err := r.cache.GetBySubdomain(ctx, subdomain)
	if err == nil {
		if err := r.validateAndAsyncCache(tenant, false); err != nil {
			return nil, err
		}
		return tenant, nil
	}

	tenant, err = r.db.GetBySubdomain(ctx, subdomain)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", baseentitieserr.ErrTenantNotFound, subdomain)
	}
	if err := r.validateAndAsyncCache(tenant, true); err != nil {
		return nil, err
	}

	return tenant, nil
}

func NewTenantResolver(cache abstractions.TenantRepository, db abstractions.TenantRepository) *TenantResolver {
	return &TenantResolver{cache, db}
}
