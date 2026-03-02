package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/abstractions/services"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/enum"
	baseentitieserr "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/error"
	asignamultitenancy "github.com/samforedev/asignads/lib/asigna-multitenancy"
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

func (r *redisService) GetById(ctx context.Context, id string) (*domain.Tenant, error) {
	key := fmt.Sprintf("tenant:id:%s", id)

	data, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, baseentitieserr.ErrTenantNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("redis get by id error: %w", err)
	}

	var tenant domain.Tenant
	if err := json.Unmarshal([]byte(data), &tenant); err != nil {
		return nil, fmt.Errorf("deserialize tenant error: %w", err)
	}

	return &tenant, nil
}

func (r *redisService) UpdateStatus(ctx context.Context, id string, status enum.TenantStatus) error {
	tenant, err := r.GetById(ctx, id)
	if err != nil {
		return nil
	}
	return r.Invalidate(ctx, tenant)
}

func (r *redisService) GetBySubDomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	key := fmt.Sprintf("tenant:%s", subdomain)

	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, baseentitieserr.ErrTenantNotFound
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
	subKey := fmt.Sprintf("tenant:%s", tenant.Subdomain)
	idKey := fmt.Sprintf("tenant:id:%s", tenant.ID)

	dsnKey := asignamultitenancy.GetTenantDSNKey(tenant.ID)

	data, _ := json.Marshal(tenant)

	pipe := r.client.Pipeline()
	pipe.Set(ctx, subKey, data, r.ttl)
	pipe.Set(ctx, idKey, data, r.ttl)
	pipe.Set(ctx, dsnKey, tenant.DBDSN, r.ttl)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *redisService) Invalidate(ctx context.Context, tenant *domain.Tenant) error {
	subdomainKey := fmt.Sprintf("tenant:%s", tenant.Subdomain)
	idKey := fmt.Sprintf("tenant:id:%s", tenant.ID)
	dsnKey := asignamultitenancy.GetTenantDSNKey(tenant.ID)
	return r.client.Del(ctx, subdomainKey, idKey, dsnKey).Err()
}
