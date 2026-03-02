package mocks

import (
	"context"
	"testing"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/business"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/enum"
	"github.com/stretchr/testify/assert"
)

func TestTenantManager_ChangeStatus(t *testing.T) {
	ctx := context.Background()
	tenantID := "uuid-123"
	newStatus := enum.INACTIVE

	mockDB := new(MockTenantRepository)
	mockCache := new(MockTenantRepository)
	manager := business.NewTenantManager(mockDB, mockCache)

	tenant := &domain.Tenant{ID: tenantID, Subdomain: "test"}

	mockDB.On("GetById", ctx, tenantID).Return(tenant, nil)
	mockDB.On("UpdateStatus", ctx, tenantID, newStatus).Return(nil)
	mockCache.On("UpdateStatus", ctx, tenantID, newStatus).Return(nil)

	err := manager.ChangeStatus(ctx, tenantID, newStatus)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}
