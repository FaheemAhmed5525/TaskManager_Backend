package handlers

import (
	"encoding/json"
	"net/http"
	"task_API/internal/models"
	"task_API/internal/storage"
	"task_API/pkg/utils"
)

type AuthHandler struct {
	storage *storage.PostgresStorage
}

func NewAuthHandler(storage *storage.PostgresStorage) *AuthHandler {
	return &AuthHandler{storage: storage}
}

func (handler AuthHandler) Register(writer http.ResponseWriter, request *http.Request) {
	var registerRequest models.RegisterUserRequest

	if error := json.NewDecoder(request.Body).Decode(&registerRequest); error != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	if registerRequest.Email == "" || registerRequest.Name == "" || registerRequest.Password == "" {
		http.Error(writer, "Some of Required Attributes empty. Email, Password, and Email can't be empty", http.StatusBadRequest)
		return
	}

	if len(registerRequest.Password) < 8 {
		http.Error(writer, "Password must be at least 8 character long", http.StatusBadRequest)
		return
	}

	// Creating password hash
	hashedPass, error := utils.HashPassword(registerRequest.Password)
	if error != nil {
		http.Error(writer, "Unable to process password", http.StatusBadRequest)
		return
	}

	// Create User
	user, error := handler.storage.CreateUser(registerRequest.Name, registerRequest.Email, hashedPass)
	if error != nil {
		http.Error(writer, "Unable to create user: "+error.Error(), http.StatusBadRequest)
		return
	}

	// Generating Token
	token, error := utils.GenerateJwtToken(user.ID, user.Email)
	if error != nil {
		http.Error(writer, "Unable to generate token", http.StatusInternalServerError)
		return
	}

	response := models.AuthResponse{
		User: models.User{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token: token,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest models.LoginUserRequest

	if error := json.NewDecoder(request.Body).Decode(&loginRequest); error != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, error := handler.storage.GetUserByEmail(loginRequest.Email)
	if error != nil {
		http.Error(writer, "Invalid email", http.StatusNotFound)
		return
	}

	if !utils.CheckHashedPassword(user.PasswordHash, loginRequest.Password) {
		http.Error(writer, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, error := utils.GenerateJwtToken(user.ID, user.Email)
	if error != nil {
		http.Error(writer, "Unable to generate token", http.StatusInternalServerError)
		return
	}

	response := models.AuthResponse{
		User: models.User{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token: token,
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

// Get Profile returns the current User Profile
func (handler *AuthHandler) GetProfile(writer http.ResponseWriter, request *http.Request) {
	// Get user ID from context( AuthMiddleware sets it)
	user, ok := request.Context().Value("user").(models.User)
	if !ok {
		http.Error(writer, "Unable to get user from context", http.StatusInternalServerError)
		return
	}

	user.PasswordHash = ""

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}
