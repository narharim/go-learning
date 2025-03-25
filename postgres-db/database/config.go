package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type dbconfig struct {
	host     string
	port     string
	dbname   string
	username string
	password string
}

func NewConfig() *dbconfig {
	return &dbconfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbname:   os.Getenv("DB_DATABASE"),
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
	}
}

func (c *dbconfig) Validate() error {
	if c.host == "" {
		return fmt.Errorf("database host is required")
	}
	if c.port == "" {
		return fmt.Errorf("database port is required")
	}
	if c.dbname == "" {
		return fmt.Errorf("database name is required")
	}
	if c.username == "" {
		return fmt.Errorf("database username is required")
	}
	if c.password == "" {
		return fmt.Errorf("database password is required")
	}
	return nil
}

func NewDB(cfg *dbconfig) (*sql.DB, error) {
	dbSource := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.username,
		cfg.password,
		cfg.host,
		cfg.port,
		cfg.dbname,
	)
	log.Println(dbSource)

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
