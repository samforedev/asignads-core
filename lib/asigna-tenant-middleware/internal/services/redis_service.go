package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/samforedev/asignads/lib/asigna-base-entities/tenant/domain"
	baseentitieserr "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/error"
	asignamultitenancy "github.com/samforedev/asignads/lib/asigna-multitenancy"
	"github.com/samforedev/asignads/lib/asigna-tenant-middleware/internal/abstractions"
)

type redisService struct {
	client *redis.Client
	ttl    time.Duration
}

func (r redisService) GetBySubdomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	key := fmt.Sprintf("tenant:%s", subdomain)
	val, err := r.client.Get(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return nil, baseentitieserr.ErrTenantNotFound
	} else if err != nil {
		return nil, fmt.Errorf("redis central db error: %w", err)
	}

	var tenant domain.Tenant
	if err := json.Unmarshal([]byte(val), &tenant); err != nil {
		return nil, fmt.Errorf("redis deserialize tenant error: %w", err)
	}

	return &tenant, nil
}

func (r redisService) SaveInCache(ctx context.Context, tenant *domain.Tenant) error {
	subKey := fmt.Sprintf("tenant:%s", tenant.Subdomain)
	dsnKey := asignamultitenancy.GetTenantDSNKey(tenant.ID)

	serializeTenant, err := json.Marshal(tenant)
	if err != nil {
		return fmt.Errorf("serialize tenant error: %w", err)
	}

	pipe := r.client.Pipeline()
	pipe.Set(ctx, subKey, serializeTenant, r.ttl)

	tenantDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		tenant.Database.Host,
		tenant.Database.Port,
		tenant.Database.User,
		tenant.Database.Password,
		tenant.Database.SslMode,
	)

	pipe.Set(ctx, dsnKey, tenantDsn, r.ttl)

	_, err = pipe.Exec(ctx)
	return err
}

func NewRedisService(client *redis.Client) abstractions.TenantRepository {
	return &redisService{
		client: client,
		ttl:    24 * time.Hour,
	}
}
