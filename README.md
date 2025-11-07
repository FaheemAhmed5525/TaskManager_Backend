# Task API - RESTful Go Server

A complete REST API built with Go for task management.

## Features

- ✅ RESTful endpoints
- ✅ JSON request/response handling
- ✅ In-memory storage with thread safety
- ✅ Request logging middleware
- ✅ Proper project structure
- ✅ Error handling

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/tasks` | Get all tasks |
| GET | `/tasks/{id}` | Get specific task |
| POST | `/tasks` | Create new task |

## How to Run

```bash
# Start the server
go run cmd/api/main.go

# Test endpoints
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl -X POST -H "Content-Type: application/json" -d '{"title": "My task"}' http://localhost:8080/tasks
