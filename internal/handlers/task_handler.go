package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task_API/internal/models"
	"task_API/internal/storage"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	storage storage.Storage // interface, not concrete type
}

func NewTaskHandler(s storage.Storage) *TaskHandler {
	return &TaskHandler{storage: s}
}

func (h *TaskHandler) GetAllTasks(writer http.ResponseWriter, request *http.Request) {
	// Get User From Context
	user, ok := request.Context().Value("user").(models.User)
	if !ok {
		http.Error(writer, "User not found in context", http.StatusInternalServerError)
		return
	}

	tasks, err := h.storage.GetAllTasks(user.ID)
	if err != nil {
		http.Error(writer, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (h *TaskHandler) GetTask(writer http.ResponseWriter, request *http.Request) {
	// Get User From Context
	user, ok := request.Context().Value("user").(models.User)
	if !ok {
		http.Error(writer, "User not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.storage.GetTaskById(id, user.ID)
	if err != nil {
		http.Error(writer, "Task not found", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(task)
}

func (h *TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	// Get User From Context
	user, ok := request.Context().Value("user").(models.User)
	if !ok {
		http.Error(writer, "User not found in context", http.StatusInternalServerError)
		return
	}

	var req models.CreateTaskRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task, err := h.storage.CreateTask(req.Title, user.ID)
	if err != nil {
		http.Error(writer, "Error creating task", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(task)
}

func (h *TaskHandler) UpdateTask(writer http.ResponseWriter, request *http.Request) {
	// Get User From Context
	user, ok := request.Context().Value("user").(models.User)
	if !ok {
		http.Error(writer, "User not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task, err := h.storage.UpdateTask(id, req.Title, req.Completed, user.ID)
	if err != nil {
		http.Error(writer, "Task not found", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(task)
}

func (h *TaskHandler) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	// Get User From Context
	user, ok := request.Context().Value("user").(models.User)
	if !ok {
		http.Error(writer, "User not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.storage.DeleteTask(id, user.ID); err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
