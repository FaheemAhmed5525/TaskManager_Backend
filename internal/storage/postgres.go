package storage

import (
	"context"
	"database/sql"
	"fmt"
	"task_API/internal/models"
	"time"
	// _ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresStorage struct {
	database *sql.DB
}

func NewPostgresStorage(conStr string) (*PostgresStorage, error) {
	db, error := sql.Open("pgx", conStr)
	if error != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", error)
	}

	// Connection check
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresStorage{database: db}, nil
}

func (storage *PostgresStorage) Close() error {
	return storage.database.Close()
}

// Create Task database interface
func (storage *PostgresStorage) CreateTask(title string) (models.Task, error) {
	query := `
	INSERT INTO tasks (title, completed, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id
	          VALUES ($1, $2, $3, $4)
			  RETURNING id, title, completed, created_at, updated_at
			  `

	var task models.Task
	now := time.Now()

	error := storage.database.QueryRow(
		query,
		title,
		false,
		now,
		now,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if error != nil {
		return models.Task{}, fmt.Errorf("failed to create task: %w", error)
	}

	return task, nil
}

// Get All Tasks database interface
func (storage *PostgresStorage) GetAllTasks() ([]models.Task, error) {
	query := `
	SELECCT id, title, completed, created_at, updated_at
	FROM tasks
	ORDER BY created_at DESC
	`

	rows, error := storage.database.Query(query)

	if error != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", error)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Get TAsk by id
func (storage *PostgresStorage) GetTaskById(id int) (models.Task, error) {
	query := `
	SELECT id, title, completed, created_at, updated_at
	FROM tasks
	WHERE id = $1
	`

	var task models.Task

	if error := storage.database.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	); error != nil {
		if error == sql.ErrNoRows {
			return models.Task{}, fmt.Errorf("task not found")
		} else if error != nil {
			return models.Task{}, fmt.Errorf("failed to get task: %w", error)
		}
	}

	return task, nil

}

func (storage *PostgresStorage) UpdateTask(id int, title string, completed bool) (models.Task, error) {
	query := `
	UPDATE tasks,
	SET title = $1, completed = $2, updated_at = $3
	WHERE id = $4
	RETURNING id, title, completed, created_at, updated_at
	`

	var task models.Task
	now := time.Now()

	error := storage.database.QueryRow(
		query,
		title,
		completed,
		now,
		id,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if error != nil {
		if error == sql.ErrNoRows {
			return models.Task{}, fmt.Errorf("task not found")
		} else {
			return models.Task{}, fmt.Errorf("failed to update task: %w", error)
		}
	}

	return task, nil
}

func (storage *PostgresStorage) DeleteTask(id int) error {
	query := `
	DELETE FROM tasks
	WHERE id = $1
	`
	result, error := storage.database.Exec(query, id)

	if error != nil {
		return fmt.Errorf("failed to delete task: %w", error)
	}

	rowsAffected, error := result.RowsAffected()
	if error != nil {
		return fmt.Errorf("failed to get rows affected: %w", error)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
