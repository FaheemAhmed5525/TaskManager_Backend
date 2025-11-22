package handlers

import (
	"encoding/json"
	"net/http"
	"task_API/internal/models"
	"task_API/internal/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (handler *AuthHandler) Register(writer http.ResponseWriter, request *http.Request) {
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

	// Create User
	authResponse, error := handler.authService.Register(&registerRequest)
	if error != nil {
		http.Error(writer, "Unable to create user: "+error.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(authResponse)
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest models.LoginUserRequest

	if error := json.NewDecoder(request.Body).Decode(&loginRequest); error != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	authResponse, error := handler.authService.Login(&loginRequest)

	if error != nil {
		http.Error(writer, error.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(authResponse)
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
