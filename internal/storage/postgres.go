package storage

import (
	"context"
	"database/sql"
	"fmt"
	"task_API/internal/models"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
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

// Create Task database interface
func (storage *PostgresStorage) CreateTask(title string, userId int) (models.Task, error) {
	query := `
	INSERT INTO tasks (title, completed, user_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, title, completed, user_id, created_at, updated_at
	`

	var task models.Task
	now := time.Now()

	error := storage.database.QueryRow(
		query,
		title,
		false,
		userId,
		now,
		now,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.UserId,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if error != nil {
		return models.Task{}, fmt.Errorf("failed to create task: %w", error)
	}

	return task, nil
}

// Get All Tasks database interface
func (storage *PostgresStorage) GetAllTasks(userId int) ([]models.Task, error) {
	query := `
	SELECT id, title, completed, user_id, created_at, updated_at
	FROM tasks
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, error := storage.database.Query(query, userId)

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
			&task.UserId,
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
func (storage *PostgresStorage) GetTaskById(id int, userId int) (models.Task, error) {
	query := `
	SELECT id, title, completed, user_id, created_at, updated_at
	FROM tasks
	WHERE id = $1 AND user_id = $2
	`

	var task models.Task

	if error := storage.database.QueryRow(query, id, userId).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.UserId,
		&task.CreatedAt,
		&task.UpdatedAt,
	); error != nil {
		if error == sql.ErrNoRows {
			return models.Task{}, fmt.Errorf("task not found")
		} else {
			return models.Task{}, fmt.Errorf("failed to get task: %w", error)
		}
	}

	return task, nil

}

func (storage *PostgresStorage) UpdateTask(id int, title string, completed bool, userId int) (models.Task, error) {
	query := `
	UPDATE tasks
	SET title = $1, completed = $2, updated_at = $3
	WHERE id = $4 AND user_id = $5
	RETURNING id, title, completed, user_id, created_at, updated_at
	`

	var task models.Task
	now := time.Now()

	error := storage.database.QueryRow(
		query,
		title,
		completed,
		now,
		id,
		userId,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.UserId,
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

func (storage *PostgresStorage) DeleteTask(id int, userId int) error {
	query := `
	DELETE FROM tasks
	WHERE id = $1 AND user_id = $2
	`
	result, error := storage.database.Exec(query, id, userId)

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

// Create User
func (storage *PostgresStorage) CreateUser(name string, email string, passwordHash string) (models.User, error) {
	query := `INSERT INTO users (name, password_hash, email, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, email, password_hash, name, created_at, updated_at`

	var user models.User
	now := time.Now()

	error := storage.database.QueryRow(
		query,
		name,
		passwordHash,
		email,
		now,
		now,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if error != nil {
		return models.User{}, fmt.Errorf("faild to create user %w", error)
	}

	return user, nil
}

func (storage *PostgresStorage) GetUserByEmail(email string) (models.User, error) {
	query := `
	SELECT id, email, password_hash, name, created_at, updated_at
	FROM users	
	WHERE email = $1
	`

	var user models.User

	err := storage.database.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		} else {
			return models.User{}, fmt.Errorf("unable to get user: %w", err)
		}
	}

	return user, nil
}

func (storage *PostgresStorage) GetUserById(id int) (models.User, error) {
	query := `SELECT id, email, password_hash, name, created_at, updated_at
	FROM users
	WHERE id = $1`

	var user models.User

	error := storage.database.QueryRow(
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if error != nil {
		if error == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		} else {
			return models.User{}, fmt.Errorf("unable to get user: %w", error)
		}
	}

	return user, nil
}
