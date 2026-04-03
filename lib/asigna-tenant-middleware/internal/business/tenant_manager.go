package business

import "github.com/samforedev/asignads/lib/asigna-tenant-middleware/internal/abstractions"

type TenantManager struct {
	cacheRepo abstractions.TenantRepository
	dbRepo    abstractions.TenantRepository
}

func NewTenantManager(cacheRepo abstractions.TenantRepository, dbRepo abstractions.TenantRepository) *TenantManager {
	return &TenantManager{cacheRepo, dbRepo}
}
