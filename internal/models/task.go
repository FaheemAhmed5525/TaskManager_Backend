package models

import "time"

// Model: Task, respsent the task entity
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"creation_time"`
	UpdatedAt time.Time `json:"updatation_time"`
}

// Model: CreateTaskRequest, represent the data required to create a task
type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	UpdatedAt time.Time `json:"updated_at"`
}
