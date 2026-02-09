package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	RedisAddr     string
	RedisPassword string
	CentralDBUrl  string
	Environment   string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:          getEnv("APP_PORT", "8080"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		CentralDBUrl:  getEnv("CENTRAL_DB_URL", "postgres://user:pass@localhost:5432/asigna_central?sslmode=disable"),
		Environment:   getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
