package main

import (
	"log"
	"net/http"
	"task_API/internal/config"
	"task_API/internal/handlers"
	"task_API/internal/storage"
	"task_API/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// choose backend:
	// store := storage.NewMemoryStorage()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database storage
	dbStorage, err := storage.NewPostgresStorage(config.GetDBConnectionString(cfg))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbStorage.Close()

	taskHandler := handlers.NewTaskHandler(dbStorage)

	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "healthy", "service": "task-api", "database": "connected"}\n`))
	}).Methods("GET")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
