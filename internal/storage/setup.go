package storage

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (storage *PostgresStorage) CreateTables() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	statements := []string{
		// User table
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			name VARCHAR(100) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMEsTAMP
		)`,

		// Task table with user relationship
		`CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			user_id INTEGER REFERENCES users(id) on DELETE CASCADE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,

		/// Indices
		`CREATE INDEX IF NOT EXISTS idx_user_email on users(email)`,
		`CREATE INDEX IF NOT EXISTS idx_task_user_id on tasks(user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_created_at on tasks(created_at)`,

		`INSERT INTO users (email, password_hash, name)
		VALUES ('faheemahmed@golang.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Faheem Ahmed')
		ON CONFLICT (email) DO NOTHING`,
	}

	for index, sttemetns := range statements {
		_, err := storage.DB.ExecContext(ctx, sttemetns)
		if err != nil {
			return fmt.Errorf("failed to execute statement %d: %w", index, err)
		}

		log.Printf("Successfully executed statement at %d", index)
	}
	log.Printf("Database created Successfully")
	return nil
}

func (storage *PostgresStorage) Close() error {
	return storage.DB.Close()
}
