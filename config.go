package main

import (
	"errors"
	"os"
)

// Config holds the runtime configuration for the service.
type Config struct {
	Env        string
	DBHost     string
	DBPassword string
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func loadConfig() (Config, error) {
	cfg := Config{
		Env:        getenv("ENV", "development"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}

	if cfg.DBHost == "" {
		return cfg, errors.New("DB_HOST is required")
	}

	return cfg, nil
}
