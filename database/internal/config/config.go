package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(envVar string) (string, error) {
	envFile := ".env.development"
	if err := godotenv.Load(envFile); err != nil {
		fmt.Printf(".env.development file not found, proceeding with environment variables\n")
	}

	value := os.Getenv(envVar)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", envVar)
	}

	return value, nil
}
