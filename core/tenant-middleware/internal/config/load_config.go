package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	RedisAddr        string
	RedisPassword    string
	CentralDBUrl     string
	Environment      string
	BackendTargetURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Port:             getEnv("APP_PORT", "8080"),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		CentralDBUrl:     getEnv("CENTRAL_DB_URL", "postgres://user:pass@localhost:5432/asigna_central?sslmode=disable"),
		Environment:      getEnv("ENVIRONMENT", "development"),
		BackendTargetURL: getEnv("BACKEND_TARGET_URL", "http://localhost:8081"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
