package storage

import (
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

func (memoryStorage *MemoryStorage) CreateTask(title string) models.Task {
	memoryStorage.mutex.Lock()
	defer memoryStorage.mutex.Unlock()

	task := models.Task{
		ID:        memoryStorage.nextId,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	memoryStorage.tasks[memoryStorage.nextId] = task
	memoryStorage.nextId++

	return task
}

func (memoryStorage *MemoryStorage) GetAllTasks() []models.Task {
	memoryStorage.mutex.RLock()
	defer memoryStorage.mutex.RUnlock()

	tasks := make([]models.Task, 0, len(memoryStorage.tasks))

	for _, task := range memoryStorage.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (memoryStorage *MemoryStorage) GetTaskById(id int) (models.Task, bool) {
	memoryStorage.mutex.RLock()
	defer memoryStorage.mutex.RUnlock()

	task, exists := memoryStorage.tasks[id]
	return task, exists
}
