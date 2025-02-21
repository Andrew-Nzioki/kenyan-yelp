package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress    string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	DatabaseURL   string
	JWTSecret     string
	JWTExpiration time.Duration
}

func LoadEvironmentVariables() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found: %v\n", err)
	}

	configuration := &Config{
		ServerAddress:    getEnvOrDefault("SERVER_ADDR", ":8080"),
		ReadTimeout:   getDurationOrDefault("READ_TIMEOUT", 15*time.Second),
		WriteTimeout:  getDurationOrDefault("WRITE_TIMEOUT", 15*time.Second),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		JWTSecret:     getEnvOrDefault("JWT_SECRET", "your-secret-key"),
		JWTExpiration: getDurationOrDefault("JWT_EXPIRATION", 24*time.Hour),
	}

	if configuration.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}else if configuration.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}else if configuration.JWTExpiration == 0 {
		return nil, fmt.Errorf("JWT_EXPIRATION is required")

	}
	return configuration, nil
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