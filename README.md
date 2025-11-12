# Task API - PostgreSQL Backend

A complete REST API with PostgreSQL database integration.

# New Features Added

- ✔️ PostgreSQL database integration
- ✔️ Complete CRUD operations (Create, Read, Update, Delete)
- ✔️ Dockerized development environment
- ✔️ Environment configuration
- ✔️ Database migrations
- ✔️ Proper error handling

# API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/tasks` | Get all tasks |
| GET | `/tasks/{id}` | Get specific task |
| POST | `/tasks` | Create new task |
| PUT | `/tasks/{id}` | Update task (partial updates supported) |
| DELETE | `/tasks/{id}` | Delete task |

# Development Setup


# Start database
docker-compose up -d

# Run application
go run cmd/api/main.go

 Access pgAdmin (database GUI)
 http://localhost:8081
 Email: admin@taskapi.com
 Password: admin_password


# Environment Variables
| Variable	|   Default  	| Description |
|-----------|-------------|-------------|
|DB_HOST	|localhost|	|Database host|
|DB_PORT	|5432	|Database port|
|DB_USER	|task_user|	Database user|
|DB_PASSWORD	|task_password|	Database password|
|DB_NAME	|task_database|	Database name|
|SERVER_PORT	|8080|	Application port|
|Testing |CRUD| Operations|


# Create
curl -X POST -H "Content-Type: application/json" -d '{"title": "New Task"}' http://localhost:8080/tasks

# Read
curl http://localhost:8080/tasks
curl http://localhost:8080/tasks/1

# Update (partial)
curl -X PUT -H "Content-Type: application/json" -d '{"completed": true}' http://localhost:8080/tasks/1

# Delete
curl -X DELETE http://localhost:8080/tasks/1

# Learning Objectives Achieved
- PostgreSQL integration with Go
- Database migrations
- Complete CRUD operations
- Environment configuration management
- Docker for development environment
- SQL query building and execution
- Proper error handling with database
