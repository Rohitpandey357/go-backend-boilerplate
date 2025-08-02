package config

import (
	"log"
	"os"
)

func GetEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("Env %s not found, using fallback: %s", key, fallback)
	return fallback
}
