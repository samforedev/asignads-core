package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/domain"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/services"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types/enums"
)

type postgresService struct {
	db *sql.DB
}

func NewPostgresService(db *sql.DB) services.TenantRepository {
	return &postgresService{db: db}
}

func (p postgresService) GetById(ctx context.Context, id string) (*domain.Tenant, error) {
	query := types.SearchTenantById

	var t domain.Tenant
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&t.ID, &t.Name, &t.Subdomain, &t.DBDSN, &t.Status, &t.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, types.ErrTenantNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("search db central error: %w", err)
	}

	return &t, nil
}

func (p postgresService) UpdateStatus(ctx context.Context, id string, status enums.TenantStatus) error {
	query := types.UpdateTenantStatus

	result, err := p.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("error updating tenant status: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return types.ErrTenantNotFound
	}

	return nil
}

func (p postgresService) GetBySubDomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	query := types.SearchTenantsQuery

	var t domain.Tenant
	err := p.db.QueryRowContext(ctx, query, subdomain).Scan(
		&t.ID, &t.Name, &t.Subdomain, &t.DBDSN, &t.Status, &t.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, types.ErrTenantNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("search db central error: %w", err)
	}

	return &t, nil
}

func (p postgresService) SaveInCache(ctx context.Context, tenant *domain.Tenant) error {
	return nil
}
