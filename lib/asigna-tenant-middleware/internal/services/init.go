package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/samforedev/asignads/lib/asigna-tenant-middleware/internal/config"
)

func InitPostgres(cfg *config.Config) *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.SslMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Connects Postgres database failed: %v", err)
	}
	return db
}

func InitRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
	})
}
