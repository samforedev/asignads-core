package asigna_multitenancy

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sony/gobreaker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	baseentitieserror "github.com/samforedev/asignads/lib/asigna-base-entities/tenant/error"
)

type TenantConnector struct {
	registry    *ConnectionRegistry
	redisClient *redis.Client
	cb          *gobreaker.CircuitBreaker
}

func NewTenantConnector(redisClient *redis.Client) *TenantConnector {

	settings := gobreaker.Settings{
		Name:        "Redis-DSN-Resolver",
		MaxRequests: 2,
		Interval:    5 * time.Second,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("Circuit Breaker [%s]: %s -> %s\n", name, from, to)
		},
	}

	return &TenantConnector{
		registry:    NewConnectionRegistry(),
		redisClient: redisClient,
		cb:          gobreaker.NewCircuitBreaker(settings),
	}
}

func (c *TenantConnector) GetDB(ctx context.Context) (*gorm.DB, error) {
	tenantID := ExtractTenantID(ctx)
	if tenantID == "" {
		return nil, baseentitieserror.ErrMissingTenantContext
	}

	db, _, exists := c.registry.Get(tenantID)
	if exists && db != nil {
		return db, nil
	}

	result, err := c.cb.Execute(func() (interface{}, error) {
		dsnKey := "tenant:" + tenantID + ":dsn"
		dsn, err := c.redisClient.Get(ctx, dsnKey).Result()
		if err != nil {
			return nil, err
		}
		return dsn, nil
	})

	if err != nil {
		return nil, fmt.Errorf("dsn resolution failed (breaker): %w", err)
	}

	dsn := result.(string)
	newDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, baseentitieserror.ErrConnectionFailed
	}

	c.registry.Set(tenantID, dsn, newDb)
	return newDb, nil
}
