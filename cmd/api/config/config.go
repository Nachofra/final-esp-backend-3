package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

const (
	host         = "localhost"
	port         = "8080"
	databaseHost = "localhost"
	databasePort = "3307"
	user         = "root"
	password     = "root"
	schema       = "clinic"
	charset      = "utf8"
)

// Config centralizes all the config of dependencies of the whole app.
type Config struct {
	DBHost      string `env:"DATABASE_HOST"`
	DBPort      string `env:"DATABASE_PORT"`
	DBUser      string `env:"DATABASE_USER"`
	DBPassword  string `env:"DATABASE_PASSWORD"`
	DBSchema    string `env:"DATABASE_SCHEMA"`
	DBCharset   string `env:"DATABASE_CHARSET"`
	DBParseTime bool   `env:"DATABASE_PARSE_TIME" envDefault:"true"`

	Host    string `env:"HOST"`
	Port    string `env:"PORT"`
	GinMode string `env:"GIN_MODE" envDefault:"debug"`
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
	cfg.DBHost = defaultValue(cfg.DBHost, databaseHost)
	cfg.DBPort = defaultValue(cfg.DBPort, databasePort)
	cfg.DBUser = defaultValue(cfg.DBUser, user)
	cfg.DBPassword = defaultValue(cfg.DBPassword, password)
	cfg.DBSchema = defaultValue(cfg.DBSchema, schema)
	cfg.DBCharset = defaultValue(cfg.DBCharset, charset)

	cfg.Host = defaultValue(cfg.Host, host)
	cfg.Port = defaultValue(cfg.Port, port)
}

// defaultValue returns the value of v if it's not empty, otherwise returns the default value.
func defaultValue(v string, def string) string {
	if v == "" {
		return def
	}

	return v
}
