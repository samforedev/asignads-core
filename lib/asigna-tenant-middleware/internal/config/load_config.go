package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	RedisAddr      string
	RedisPassword  string
	DbHost         string
	DbPort         string
	DbUser         string
	DbPassword     string
	DbName         string
	SslMode        string
	AppEnvironment string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		AppPort:        getEnv("APP_PORT", "8080"),
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:  getEnv("REDIS_PASSWORD", ""),
		DbHost:         getEnv("DB_HOST", "localhost"),
		DbPort:         getEnv("DB_PORT", "5432"),
		DbUser:         getEnv("DB_USER", "postgres"),
		DbPassword:     getEnv("DB_PASSWORD", ""),
		DbName:         getEnv("DB_NAME", "postgres"),
		SslMode:        getEnv("SSL_MODE", "disable"),
		AppEnvironment: getEnv("APP_ENVIRONMENT", "development"),
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
