package business

import (
	"context"
	"fmt"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/services"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/enum"
)

type TenantManager struct {
	dbRepo    services.TenantRepository
	cacheRepo services.TenantRepository
}

func NewTenantManager(db services.TenantRepository, cache services.TenantRepository) *TenantManager {
	return &TenantManager{
		dbRepo:    db,
		cacheRepo: cache,
	}
}

func (m *TenantManager) ChangeStatus(ctx context.Context, id string, status enum.TenantStatus) error {
	_, err := m.dbRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if err := m.dbRepo.UpdateStatus(ctx, id, status); err != nil {
		return err
	}

	if err := m.cacheRepo.UpdateStatus(ctx, id, status); err != nil {
		fmt.Printf("[WARN] Failed to invalidate cache for tenant %s: %v\n", id, err)
	}

	return nil
}
