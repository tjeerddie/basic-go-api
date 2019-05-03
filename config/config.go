package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Getenv gets ENV variables with default
func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

// ReadDotEnv populates environment variables with a dot env file.
func ReadDotEnv(configFile string) {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Fatal("Error loading .env file", configFile)
	}
}
