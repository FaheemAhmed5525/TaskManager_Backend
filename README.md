# Task API - PostgreSQL Backend with Authentication

A complete REST API with PostgreSQL database integration, JWT authentication, and user management.

## 🚀 New Features Added

### Authentication & Security
- ✅ User registration and login
- ✅ JWT token generation and validation  
- ✅ Password hashing with bcrypt
- ✅ Authentication middleware
- ✅ User-specific task management
- ✅ Protected API endpoints
- ✅ Secure password storage

### Database & Architecture
- ✅ PostgreSQL database integration
- ✅ Complete CRUD operations
- ✅ User-task relationships
- ✅ Environment configuration management
- ✅ Docker development environment
- ✅ Database migrations and setup

## 📋 API Endpoints

### Public Endpoints (No Authentication Required)
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/register` | Register new user |
| `POST` | `/login` | User login |
| `GET` | `/health` | Health check |

### Protected Endpoints (Authentication Required)
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/tasks` | Get user's tasks |
| `GET` | `/tasks/{id}` | Get specific task |
| `POST` | `/tasks` | Create new task |
| `PUT` | `/tasks/{id}` | Update task (partial updates supported) |
| `DELETE` | `/tasks/{id}` | Delete task |
| `GET` | `/profile` | Get user profile |

## 🛠️ Development Setup

### Prerequisites
- Docker and Docker Compose
- Go 1.21+

### Quick Start
# 1. Start database services
docker-compose up -d

# 2. Run the application
go run cmd/api/main.go

# 3. Access pgAdmin (database GUI)
# http://localhost:8081
# Email: admin@taskapi.com  
# Password: adminpass

Manual Database Setup (If Needed)
bash

# Run database setup script
go run scripts/setup_db.go

🔧 Environment Variables
Variable	Default	Description
DB_HOST	localhost	Database host
DB_PORT	5432	Database port
DB_USER	taskuser	Database user
DB_PASSWORD	taskpass	Database password
DB_NAME	taskdb	Database name
SERVER_PORT	8080	Application port
JWT_SECRET	your-secret-key	JWT signing key


## 🔐 Authentication Flow
# 1. Register a New User
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe"
  }'

# 2. Login to Get Token
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'

Response:
json

{
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "created_at": "2024-11-05T10:00:00Z",
    "updated_at": "2024-11-05T10:00:00Z"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

# 3. Use Token for Protected Requests

# Get user profile
curl http://localhost:8080/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Create a task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title": "My private task"}'

# Get user's tasks
curl http://localhost:8080/tasks \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

