package services

import "task_API/internal/models"

type AuthService interface {
	Register(request *models.RegisterUserRequest) (*models.AuthResponse, error)
	Login(request *models.LoginUserRequest) (*models.AuthResponse, error)
	ValidateToken(token string) (*models.User, error)
}

type TaskService interface {
	CreateTask(req *models.CreateTaskRequest, userId int) (*models.Task, error)
	GetTaskById(taskId int, userId int) (*models.Task, error)
	GetAllTasks(userId int) ([]*models.Task, error)
	UpdateTask(taskId int, req *models.UpdateTaskRequest, userId int) (*models.Task, error)
}
