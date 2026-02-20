package mocks

import (
	"context"
	"testing"
	"time"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/business"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/enum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTenantResolver_Resolve(t *testing.T) {
	ctx := context.Background()
	host := "pruebauno.localhost"
	subdomain := "pruebauno"

	t.Run("Debe retornar tenant desde cache si existe", func(t *testing.T) {
		mockDB := new(MockTenantRepository)
		mockCache := new(MockTenantRepository)
		resolver := business.NewTenantResolver(mockCache, mockDB)

		expectedTenant := &domain.Tenant{
			ID:        "uuid-123",
			Subdomain: subdomain,
			Status:    enum.ACTIVE,
		}

		mockCache.On("GetBySubDomain", ctx, subdomain).Return(expectedTenant, nil)

		result, err := resolver.Resolve(ctx, host)

		assert.NoError(t, err)
		assert.Equal(t, expectedTenant.ID, result.ID)

		mockDB.AssertNotCalled(t, "GetBySubDomain", mock.Anything, mock.Anything)
	})

	t.Run("Debe buscar en DB y guardar en cache si no está en caché", func(t *testing.T) {
		mockDB := new(MockTenantRepository)
		mockCache := new(MockTenantRepository)
		resolver := business.NewTenantResolver(mockCache, mockDB)

		tenantFromDB := &domain.Tenant{
			ID:        "uuid-from-db",
			Subdomain: subdomain,
			Status:    enum.ACTIVE,
		}

		mockCache.On("GetBySubDomain", ctx, subdomain).Return(nil, assert.AnError)
		mockDB.On("GetBySubDomain", ctx, subdomain).Return(tenantFromDB, nil)
		mockCache.On("SaveInCache", ctx, tenantFromDB).Return(nil)

		result, err := resolver.Resolve(ctx, host)
		if assert.NoError(t, err) {
			assert.NotNil(t, result)
			assert.Equal(t, "uuid-from-db", result.ID)
		}
		time.Sleep(10 * time.Millisecond)
		mockDB.AssertExpectations(t)
		mockCache.AssertExpectations(t)
	})
}
