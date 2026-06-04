package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	AppEnv       string
	HTTPAddr     string
	JWTSecret    string
	DatabasePath string
	CORSOrigins  []string
}

func Load() Config {
	return Config{
		AppEnv:       envOr("APP_ENV", "development"),
		HTTPAddr:     envOr("HTTP_ADDR", ":8080"),
		JWTSecret:    envOr("JWT_SECRET", "dev-secret-change-me"),
		DatabasePath: envOr("DATABASE_PATH", "./data/demo.db"),
		CORSOrigins:  splitCSV(envOr("CORS_ORIGINS", "http://localhost:5173")),
	}
}

func (c Config) Validate() error {
	if strings.TrimSpace(c.JWTSecret) == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	if strings.TrimSpace(c.DatabasePath) == "" {
		return fmt.Errorf("DATABASE_PATH is required")
	}
	return nil
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func splitCSV(raw string) []string {
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
