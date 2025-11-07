package main

import (
	"log"
	"net/http"
	"task_API/internal/handlers"
	"task_API/internal/storage"
	"task_API/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Storage initialization
	storage := storage.NewMemoryStorage()

	// TaskHandler initialization
	taskHandler := handlers.NewTaskHandlers(storage)

	// Route
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	// Routes
	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")

	// Health check endpoint
	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"status": "healthy", "service": "task_API"}`))
	}).Methods("GET")

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
