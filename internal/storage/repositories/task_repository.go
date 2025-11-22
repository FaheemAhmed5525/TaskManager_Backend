package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"task_API/internal/models"
	"task_API/internal/storage"
)

type taskRepository struct {
	database *sql.DB
}

func NewTaskRepository(database *storage.PostgresStorage) TaskRepository {
	return &taskRepository{database: database.DB}
}

// Create Task database interface
func (repo *taskRepository) CreateTask(t *models.Task) error {
	query := `
	INSERT INTO tasks (title, completed, user_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, title, completed, user_id, created_at, updated_at
	`

	var task models.Task

	error := repo.database.QueryRow(
		query,
		t.Title,
		t.Completed,
		t.UserId,
		t.CreatedAt,
		t.UpdatedAt,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.UserId,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if error != nil {
		return fmt.Errorf("failed to create task: %w", error)
	}

	return nil
}

// Get All Tasks database interface
func (repo *taskRepository) GetAllTasks(userId int) ([]*models.Task, error) {
	query := `
	SELECT id, title, completed, user_id, created_at, updated_at
	FROM tasks
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, error := repo.database.Query(query, userId)

	if error != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", error)
	}
	defer rows.Close()

	var tasks []*models.Task
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
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// Get TAsk by id
func (repo *taskRepository) GetTaskById(id int) (*models.Task, error) {
	query := `
	SELECT id, title, completed, user_id, created_at, updated_at
	FROM tasks
	WHERE id = $1
	`

	var task models.Task

	if error := repo.database.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.UserId,
		&task.CreatedAt,
		&task.UpdatedAt,
	); error != nil {
		if error == sql.ErrNoRows {
			return &models.Task{}, fmt.Errorf("task not found")
		} else {
			return &models.Task{}, fmt.Errorf("failed to get task: %w", error)
		}
	}

	return &task, nil

}

func (repo *taskRepository) UpdateTask(task *models.Task) error {
	query := `
	UPDATE tasks
	SET title = $1, completed = $2, updated_at = $3
	WHERE id = $4 AND user_id = $5
	RETURNING id, title, completed, user_id, created_at, updated_at
	`

	var t models.Task

	error := repo.database.QueryRow(
		query,
		task.Title,
		task.Completed,
		task.Completed,
		task.ID,
		task.UserId,
	).Scan(
		&t.ID,
		&t.Title,
		&t.Completed,
		&t.UserId,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	if error != nil {
		if error == sql.ErrNoRows {
			return fmt.Errorf("task not found")
		} else {
			return fmt.Errorf("failed to update task: %w", error)
		}
	}

	return nil
}

func (repo *taskRepository) DeleteTask(id int) error {
	query := `
	DELETE FROM tasks
	WHERE id = $1 
	`
	result, error := repo.database.Exec(query, id)

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
func (repo *taskRepository) CreateUser(name string, email string, passwordHash string) (models.User, error) {
	query := `INSERT INTO users (name, password_hash, email, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, email, password_hash, name, created_at, updated_at`

	var user models.User
	now := time.Now()

	error := repo.database.QueryRow(
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

func (repo *taskRepository) GetUserByEmail(email string) (models.User, error) {
	query := `
	SELECT id, email, password_hash, name, created_at, updated_at
	FROM users	
	WHERE email = $1
	`

	var user models.User

	err := repo.database.QueryRow(
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

func (repo *taskRepository) GetUserById(id int) (models.User, error) {
	query := `SELECT id, email, password_hash, name, created_at, updated_at
	FROM users
	WHERE id = $1`

	var user models.User

	error := repo.database.QueryRow(
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
