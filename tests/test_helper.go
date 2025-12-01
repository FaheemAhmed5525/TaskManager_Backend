package tests

import (
	"database/sql"
	"log"
	"os"
	"task_API/internal/config"
	"task_API/internal/storage"
	"task_API/internal/storage/repositories"
	"task_API/pkg/database"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB     *sql.DB
	testConfig *config.Config
	userRepo   repositories.UserRepository
	taskRepo   repositories.TaskRepository
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	// Loadng the test configs
	testConfig = getTestConfig()

	var err error
	testDB, err = database.NewTestDB(testConfig.Database)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	testStorage := storage.PostgresStorage{DB: testDB}

	userRepo = repositories.NewUserRepository(&testStorage)
	taskRepo = repositories.NewTaskRepository(&testStorage)

	// Creating tables
	createTestTables()
}

func createTestTables() {
	queries := []string{
		`DROP TABLE IF EXISTS tasks, users CASCADE`,

		`CREATE TABLE IF NOT EXIST users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			name VARCHAR(100) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,

		// Task table with user relationship
		`CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			user_id INTEGER REFERENCES users(id) on DELETE CASCADE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,

		/// Indices
		`CREATE INDEX IF NOT EXISTS idx_user_email on users(email)`,
		`CREATE INDEX IF NOT EXISTS idx_task_user_id on tasks(user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_created_at on tasks(created_at)`,
	}

	for _, query := range queries {
		_, err := testDB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to execute query: %s \nError: %v", query, err)
		}
	}
}

func tearDown() {
	if testDB != nil {
		testDB.Close()
	}
}

// Returns Test database current instance
func GetTestDB() *sql.DB {
	return testDB
}

// Returns Test configs
func GetTestConfig() *config.Config {
	return testConfig
}
