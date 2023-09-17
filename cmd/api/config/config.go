package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "root"
	schema   = "clinic"
)

// Config centralizes all the config of dependencies of the whole app.
type Config struct {
	DBHost     string `env:"DATABASE_HOST"`
	DBPort     string `env:"DATABASE_HOST"`
	DBUser     string `env:"DATABASE_USER"`
	DBPassword string `env:"DATABASE_PASSWORD"`
	DBSchema   string `env:"DATABASE_SCHEMA"`
}

// Get returns the config of the whole app.
func Get() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	loadDefaults(cfg)

	return cfg, nil
}

// loadDefaults puts default values in the configuration if some environment variables are not set.
func loadDefaults(cfg *Config) {
	cfg.DBHost = defaultValue(cfg.DBHost, host)
	cfg.DBPort = defaultValue(cfg.DBPort, port)
	cfg.DBUser = defaultValue(cfg.DBUser, user)
	cfg.DBPassword = defaultValue(cfg.DBPassword, password)
	cfg.DBSchema = defaultValue(cfg.DBSchema, schema)
}

// defaultValue returns the value of v if it's not empty, otherwise returns the default value.
func defaultValue(v string, def string) string {
	if v == "" {
		return def
	}

	return v
}
