package test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	baseentitiesconstants "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/constant"
	asignamultitenancy "github.com/samforedev/asignads/lib/asigna-multitenancy"
	"github.com/stretchr/testify/assert"
)

func TestConnector_CircuitBreaker_Open(t *testing.T) {
	badRedis := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:1234",
		DialTimeout: time.Millisecond * 10,
		ReadTimeout: time.Millisecond * 10,
		MaxRetries:  -1,
	})
	connector := asignamultitenancy.NewTenantConnector(badRedis)
	ctx := context.WithValue(context.Background(), baseentitiesconstants.TenantIDKey, "test-id")

	for i := 0; i < 3; i++ {
		_, _ = connector.GetDB(ctx)
	}

	_, err := connector.GetDB(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "circuit breaker is open")
}

func TestConnector_CircuitBreaker_Success(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	redisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	connector := asignamultitenancy.NewTenantConnector(redisClient)

	tenantID := "test-xyz"
	expectedDSN := "host=localhost user=gorm password=gorm dbname=gorm port=9920"

	key := asignamultitenancy.GetTenantDSNKey(tenantID)
	_ = mr.Set(key, expectedDSN)

	ctx := context.WithValue(context.Background(), baseentitiesconstants.TenantIDKey, tenantID)
	_, err := connector.GetDB(ctx)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to establish connection to tenant database")
}
