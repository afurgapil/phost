package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(envVar string) (string, error) {
	// We use .env.development in development phase
	envFile := ".env.development"

	if err := godotenv.Load(envFile); err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	value := os.Getenv(envVar)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", envVar)
	}

	return value, nil
}
