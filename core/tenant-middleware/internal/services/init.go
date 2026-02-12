package services

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/samforedev/asignads/core/tenant-middleware/internal/config"
	"log"
)

func InitRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
	})
}

func InitPostgres(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.CentralDBUrl)
	if err != nil {
		log.Fatalf("Connects Postgres database failed: %v", err)
	}
	return db
}
