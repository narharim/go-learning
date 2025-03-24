package database

import (
	"database/sql"
	"fmt"
	"os"
)

type DBConfig struct {
	host     string
	Port     string
	DBName   string
	Username string
	Password string
}

func NewConfig() *DBConfig {
	return &DBConfig{
		host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func (c *DBConfig) Validate() error {
	if c.host == "" {
		return fmt.Errorf("database host is required")
	}
	if c.Port == "" {
		return fmt.Errorf("database port is required")
	}
	if c.DBName == "" {
		return fmt.Errorf("database name is required")
	}
	if c.Username == "" {
		return fmt.Errorf("database username is required")
	}
	if c.Password == "" {
		return fmt.Errorf("database password is required")
	}
	return nil
}

func NewDB(cfg *DBConfig) (*sql.DB, error) {
	dbSource := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
