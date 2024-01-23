package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	JWT_SECRET string
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if fallback == "" {
		slog.Error("Environment variable " + key + " is not set")
		os.Exit(1)
	}

	return fallback
}

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found")
	}

	PORT = getEnv("PORT", "3000")
	JWT_SECRET = getEnv("JWT_SECRET", "")
}
