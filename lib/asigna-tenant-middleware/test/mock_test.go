package test

import (
	"context"

	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	"github.com/stretchr/testify/mock"
)

type MockTenantRepository struct {
	mock.Mock
}

func (m *MockTenantRepository) GetBySubdomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	args := m.Called(ctx, subdomain)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Tenant), args.Error(1)
}

func (m *MockTenantRepository) SaveInCache(ctx context.Context, tenant *domain.Tenant) error {
	args := m.Called(ctx, tenant)
	return args.Error(0)
}
