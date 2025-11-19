package repositories

import "task_API/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTaskById(id int) (*models.Task, error)
	GetAllTasks(userId int) ([]*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id int) error
}
