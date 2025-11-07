package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"task_API/internal/models"
	"task_API/internal/storage"

	"github.com/gorilla/mux"
)

type TaksHandler struct {
	storage *storage.MemoryStorage
}

func NewTaskHandlers(storage *storage.MemoryStorage) *TaksHandler {
	return &TaksHandler{storage: storage}
}

func (handler *TaksHandler) GetAllTasks(writer http.ResponseWriter, r *http.Request) {
	tasks := handler.storage.GetAllTasks()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (handler *TaksHandler) GetTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		http.Error(writer, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, exists := handler.storage.GetTaskById(id)
	if !exists {
		http.Error(writer, "Task not found", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(task)
}

func (handler *TaksHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var req models.CreateTaskRequest

	if error := json.NewDecoder(request.Body).Decode(&req); error != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(writer, "Title is required", http.StatusBadRequest)
		return
	}

	task := handler.storage.CreateTask(req.Title)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(task)
}
