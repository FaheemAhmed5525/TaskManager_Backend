package storage

import (
	"fmt"
	"sync"
	"task_API/internal/models"
	"time"
)

type MemoryStorage struct {
	tasks  map[int]models.Task
	mutex  sync.RWMutex
	nextId int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int]models.Task),
		nextId: 1,
	}
}

func (m *MemoryStorage) CreateTask(title string) (models.Task, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	task := models.Task{
		ID:        m.nextId,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.tasks[m.nextId] = task
	m.nextId++
	return task, nil
}

func (m *MemoryStorage) GetAllTasks() ([]models.Task, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	tasks := make([]models.Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (m *MemoryStorage) GetTaskById(id int) (models.Task, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	task, exists := m.tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("task not found")
	}
	return task, nil
}

func (m *MemoryStorage) UpdateTask(id int, title string, completed bool) (models.Task, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	task, exists := m.tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("task not found")
	}

	task.Title = title
	task.Completed = completed
	task.UpdatedAt = time.Now()
	m.tasks[id] = task

	return task, nil
}

func (m *MemoryStorage) DeleteTask(id int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.tasks[id]; !exists {
		return fmt.Errorf("task not found")
	}
	delete(m.tasks, id)
	return nil
}
