package repositories

import (
	"database/sql"
	"fmt"

	"task_API/internal/models"
)

type userRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{database: database}
}

func (repo *userRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users(name, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	error := repo.database.QueryRow(
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if error != nil {
		return fmt.Errorf("failed to create user: %v", error)
	}

	return nil
}

func (repo *userRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := repo.database.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		} else {
			return nil, fmt.Errorf("failed to get user by email: %v", err)
		}
	}
	return &user, nil
}

func (repo *userRepository) GetUserById(id int) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := repo.database.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		} else {
			return nil, fmt.Errorf("failed to get user by id: %v", err)
		}
	}
	return &user, nil
}

func (repo *userRepository) UpdateUser(user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, password_hash = $3, updated_at = $4
		WHERE id = $5
		RETURNING updated_at
	`

	error := repo.database.QueryRow(
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.UpdatedAt,
		user.ID,
	).Scan(
		&user.UpdatedAt,
	)

	if error != nil {
		return fmt.Errorf("failed to update user: %v", error)
	}

	return nil
}

func (repo *userRepository) DeleteUser(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	error := repo.database.QueryRow(query, id).Err()
	if error != nil {
		return fmt.Errorf("failed to delete user: %v", error)
	}
	return nil
}
