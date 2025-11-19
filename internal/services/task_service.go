package services

import (
	"fmt"
	"time"

	"task_API/internal/models"
	"task_API/internal/storage/repositories"
)

type taskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) *taskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (tService *taskService) CreateTask(req *models.CreateTaskRequest, userId int) (*models.Task, error) {
	task := &models.Task{
		Title:     req.Title,
		Completed: false,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := tService.taskRepo.CreateTask(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (tService *taskService) GetTaskById(taskId int, userId int) (*models.Task, error) {
	task, err := tService.taskRepo.GetTaskById(taskId)
	if err != nil {
		return nil, err
	}
	if task.UserId != userId {
		return nil, fmt.Errorf("unauthorized access to task")
	}
	return task, nil
}

func (tService *taskService) GetAllTasks(userId int) ([]*models.Task, error) {
	tasks, err := tService.taskRepo.GetAllTasks(userId)
	if err != nil {
		return nil, err
	}

	var filtered []*models.Task // must be a slice, not a single pointer

	for _, task := range tasks {
		if task.UserId == userId {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func (tService *taskService) UpdateTask(taskId int, req *models.UpdateTaskRequest, userId int) (*models.Task, error) {
	existing, err := tService.taskRepo.GetTaskById(taskId)

	if err != nil {
		return nil, err
	}
	if existing.UserId != userId {
		return nil, fmt.Errorf("resource not allowed")
	}

	task := &models.Task{
		Title:     req.Title,
		Completed: req.Completed,
		UserId:    userId,
		CreatedAt: existing.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = tService.taskRepo.UpdateTask(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
