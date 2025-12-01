package tests

import (
	"os"
	"task_API/internal/config"
	"time"

	_ "github.com/lib/pq"
)

func getTestConfig() *config.Config {
	// Seting OS sesttings
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_PASSWORD", "test_password")
	os.Setenv("DB_NAME", "task_database")
	os.Setenv("JWT_SECRET", "test_jwt_secret")
	os.Setenv("ENVIRONMENT", "test")

	config, err := config.Load()
	if err != nil {
		panic("Failed to load testign configs: " + err.Error())
	}

	return config
}

func GetTestDBConfig() *config.DatabaseConfig {
	return &config.DatabaseConfig{
		DBHost:          "localhost",
		DBPort:          5432,
		DBUser:          "task_user",
		DBPassword:      "task_password",
		DBName:          "task_database",
		MaxOpenConns:    25,
		MaxIdleConns:    25,
		ConnMaxLifetime: 5 * time.Minute,
	}
}
