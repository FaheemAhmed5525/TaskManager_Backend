package storage

import "task_API/internal/models"

type Storage interface {
	CreateTask(title string) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	UpdateTask(id int, title string, completed bool) (models.Task, error)
	DeleteTask(id int) error
}
