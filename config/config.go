package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	DBDSN string
	Port  string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Load() *ServerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	cfg := &ServerConfig{
		DBDSN: getEnv("DB_DSN", ""),
		Port:  getEnv("PORT", "8000"),
	}

	return cfg
}