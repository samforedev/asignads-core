package mocks

import (
	"context"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/domain"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types/enums"
	"github.com/stretchr/testify/mock"
)

type MockTenantRepository struct {
	mock.Mock
}

func (m *MockTenantRepository) GetById(ctx context.Context, id string) (*domain.Tenant, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Tenant), args.Error(1)
}

func (m *MockTenantRepository) GetBySubDomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	args := m.Called(ctx, subdomain)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Tenant), args.Error(1)
}

func (m *MockTenantRepository) UpdateStatus(ctx context.Context, id string, status enums.TenantStatus) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

func (m *MockTenantRepository) SaveInCache(ctx context.Context, tenant *domain.Tenant) error {
	args := m.Called(ctx, tenant)
	return args.Error(0)
}
