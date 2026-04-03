package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/constant"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	baseentitieserr "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/error"
	"github.com/samforedev/asignads/lib/asigna-tenant-middleware/internal/abstractions"
)

type postgresService struct {
	db *sql.DB
}

func (p postgresService) GetBySubdomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	query := constant.SearchTenantBySubdomain

	var tenant domain.Tenant
	err := p.db.QueryRowContext(ctx, query, subdomain).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.Subdomain,
		&tenant.Database.Host,
		&tenant.Database.Port,
		&tenant.Database.DbName,
		&tenant.Database.User,
		&tenant.Database.Password,
		&tenant.Database.SslMode,
		&tenant.StatusId,
		&tenant.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, baseentitieserr.ErrTenantNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("search central db error: %w", err)
	}

	return &tenant, nil
}

func (p postgresService) SaveInCache(ctx context.Context, tenant *domain.Tenant) error {
	return nil
}

func NewPostgresService(db *sql.DB) abstractions.TenantRepository {
	return &postgresService{
		db: db,
	}
}
