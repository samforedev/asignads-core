package services

import (
	"context"

	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/enum"
)

type TenantRepository interface {
	GetById(ctx context.Context, id string) (*domain.Tenant, error)
	UpdateStatus(ctx context.Context, id string, status enum.TenantStatus) error
	GetBySubDomain(ctx context.Context, subdomain string) (*domain.Tenant, error)
	SaveInCache(ctx context.Context, tenant *domain.Tenant) error
}
