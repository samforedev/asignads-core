package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/domain"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/services"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/types"
	asigna_multitenancy "github.com/samforedev/asignads/lib/asigna-multitenancy"
)

type redisService struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisService(client *redis.Client) services.TenantRepository {
	return &redisService{
		client: client,
		ttl:    24 * time.Hour,
	}
}

func (r *redisService) GetBySubDomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	key := fmt.Sprintf("tenant:%s", subdomain)

	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, types.ErrTenantNotFound
	} else if err != nil {
		return nil, fmt.Errorf("redis error: %w", err)
	}

	var tenant domain.Tenant
	if err := json.Unmarshal([]byte(val), &tenant); err != nil {
		return nil, fmt.Errorf("deserialize tenant error: %w", err)
	}

	return &tenant, nil
}

func (r *redisService) SaveInCache(ctx context.Context, tenant *domain.Tenant) error {
	key := fmt.Sprintf("tenant:%s", tenant.Subdomain)

	data, err := json.Marshal(tenant)
	if err != nil {
		return fmt.Errorf("serialize tenant error: %w", err)
	}

	if err := r.client.Set(ctx, key, data, r.ttl).Err(); err != nil {
		return err
	}

	dsnKey := asigna_multitenancy.GetTenantDSNKey(tenant.ID)

	return r.client.Set(ctx, dsnKey, tenant.DBDSN, r.ttl).Err()
}
