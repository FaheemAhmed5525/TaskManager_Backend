package database

import (
	"database/sql"
	"fmt"
	"task_API/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Next database testing connection
func NewTestDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	conStr := fmt.Sprintf(
		"postgress://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("pgx", conStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open test database: %w", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0) // Longer for testing

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping testing database: %w", err)
	}

	return db, nil
}
