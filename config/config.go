package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all environment variables
type Config struct {
	Port          string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       string
	URLExpTime    time.Duration
}

// LoadConfig reads environment variables from the .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found. Using default values.")
	}

	return &Config{
		Port:          getEnv("PORT", "8080"),
		RedisDB:       getEnv("REDIS_DB", "0"),
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		URLExpTime:    7 * 24 * time.Hour, // 7 days expiry
	}
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
