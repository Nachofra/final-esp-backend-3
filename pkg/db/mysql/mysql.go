package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Config centralizes the parameters required to construct the MySQL database URL.
type Config struct {
	username string
	password string
	host     string
	name     string
}

// New creates a new MySQL configuration by applying all the provided options to it.
// You can pass a series of options as variadic arguments to customize the configuration.
func New(options ...func(*Config)) *Config {
	db := &Config{}
	for _, o := range options {
		o(db)
	}
	return db
}

// Open establishes a connection to the database using the provided configuration and returns the database connection.
func Open(cfg *Config) (*sql.DB, error) {
	db, err := cfg.start()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// start opens a MySQL database connection using the current configuration.
func (cfg *Config) start() (*sql.DB, error) {
	return sql.Open("mysql", cfg.getConnectionString())
}

// getConnectionString generates the MySQL connection string based on the current configuration.
func (cfg *Config) getConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.username, cfg.password, cfg.host, cfg.name)
}

// WithUsername is used to set database username in the config.
func WithUsername(username string) func(*Config) {
	return func(db *Config) {
		db.username = username
	}
}

// WithPassword is used to set database password in the config.
func WithPassword(password string) func(*Config) {
	return func(db *Config) {
		db.password = password
	}
}

// WithHost is used to set database host in the config.
func WithHost(host string) func(*Config) {
	return func(db *Config) {
		db.host = host
	}
}

// WithName is used to set database name in the config.
func WithName(name string) func(*Config) {
	return func(db *Config) {
		db.name = name
	}
}
