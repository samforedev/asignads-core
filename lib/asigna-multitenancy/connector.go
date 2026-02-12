package asigna_multitenancy

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TenantConnector struct {
	registry    *ConnectionRegistry
	redisClient *redis.Client
}

func NewTenantConnector(redisClient *redis.Client) *TenantConnector {
	return &TenantConnector{
		registry:    NewConnectionRegistry(),
		redisClient: redisClient,
	}
}

func (c *TenantConnector) GetDB(ctx context.Context) (*gorm.DB, error) {
	id := ExtractTenantID(ctx)
	if id == "" {
		return nil, ErrMissingTenantContext
	}

	// Cache en RAM del microservicio
	db, _, exists := c.registry.Get(id)
	if exists && db != nil {
		return db, nil
	}

	// Cache en REDIS (Capa de seguridad)
	dsnKey := "tenant:" + id + ":dsn"
	dsn, err := c.redisClient.Get(ctx, dsnKey).Result()
	if err != nil {
		return nil, fmt.Errorf("could not resolve DSN from redis for tenant %s: %w", id, err)
	}

	// Conectar a la DB
	newDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, ErrConnectionFailed
	}

	// Guarda registro para la proxima conexion
	c.registry.Set(id, dsn, newDb)
	return newDb, nil
}
