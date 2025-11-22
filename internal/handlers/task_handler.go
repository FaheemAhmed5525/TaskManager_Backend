package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task_API/internal/models"
	"task_API/internal/services"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) GetAllTasks(writer http.ResponseWriter, request *http.Request) {
	// Get User From Context
	user, ok := request.Context().Value("user").(models.User)
	if !ok {
		http.Error(writer, "User not found in context", http.StatusInternalServerError)
		return
	}

	tasks, err := h.taskService.GetAllTasks(user.ID)
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

	task, err := h.taskService.GetTaskById(id, user.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
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

	task, err := h.taskService.CreateTask(&req, user.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
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

	var req models.UpdateTaskRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskService.UpdateTask(id, &req, user.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
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

	if err := h.taskService.DeleteTask(id, user.ID); err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
