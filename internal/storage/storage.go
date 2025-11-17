package storage

import "task_API/internal/models"

type Storage interface {
	// User methods
	CreateUser(name string, email string, passwordHash string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id int) (models.User, error)

	// Tasks methods
	CreateTask(title string, userID int) (models.Task, error)
	GetAllTasks(userID int) ([]models.Task, error)
	GetTaskById(id int, userID int) (models.Task, error)
	UpdateTask(id int, title string, completed bool, userID int) (models.Task, error)
	DeleteTask(id int, userID int) error

	// setup
	CreateTables() error
	Close() error
}
