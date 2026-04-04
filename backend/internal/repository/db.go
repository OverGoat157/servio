package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"servio/internal/config"
)

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("connecting to db: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	return db, nil
}

func RunMigrations(db *sqlx.DB) error {
	data, err := os.ReadFile("migrations/001_init.up.sql")
	if err != nil {
		return fmt.Errorf("reading migration file: %w", err)
	}
	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("executing migration: %w", err)
	}
	return nil
}
