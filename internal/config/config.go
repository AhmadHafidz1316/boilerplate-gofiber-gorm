package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	DBDriver string
	DBDsn    string
}

func Load() *Config {
	_ = godotenv.Load() // ignore if .env not present

	return &Config{
		AppPort:  getEnv("APP_PORT", "3000"),
		DBDriver: getEnv("DB_DRIVER", "sqlite"),
		DBDsn:    getEnv("DB_DSN", "file:app.db?cache=shared&mode=rwc"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
