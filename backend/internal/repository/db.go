package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"ab-team/internal/config"
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
	files := []string{
		"migrations/001_init.up.sql",
		"migrations/002_restaurant_cover_promo.up.sql",
		"migrations/003_user_role.up.sql",
	}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("reading migration %s: %w", f, err)
		}
		_, err = db.Exec(string(data))
		if err != nil {
			return fmt.Errorf("executing migration %s: %w", f, err)
		}
	}
	return nil
}
