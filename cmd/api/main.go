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
	// choose one backend:
	// store, _ := storage.NewPostgresStorage("postgres://user:pass@localhost:5432/db")
	store := storage.NewMemoryStorage()

	taskHandler := handlers.NewTaskHandler(store)

	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
