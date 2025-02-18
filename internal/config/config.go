package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr    string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	DatabaseURL   string
	JWTSecret     string
	JWTExpiration time.Duration
}

func Load() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found: %v\n", err)
	}

	cfg := &Config{
		ServerAddr:    getEnvOrDefault("SERVER_ADDR", ":8080"),
		ReadTimeout:   getDurationOrDefault("READ_TIMEOUT", 15*time.Second),
		WriteTimeout:  getDurationOrDefault("WRITE_TIMEOUT", 15*time.Second),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		JWTSecret:     getEnvOrDefault("JWT_SECRET", "your-secret-key"),
		JWTExpiration: getDurationOrDefault("JWT_EXPIRATION", 24*time.Hour),
	}

	// Validate required fields
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	strValue := os.Getenv(key)
	if strValue == "" {
		return defaultValue
	}
	value, err := time.ParseDuration(strValue)
	if err != nil {
		return defaultValue
	}
	return value
}