package services

import (
	"fmt"
	"time"

	"task_API/internal/models"
	"task_API/internal/storage/repositories"
	"task_API/pkg/utils"
)

type authService struct {
	userRepo      repositories.UserRepository
	jwtSecret     string
	jwtExpiration time.Duration
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string, jwtExpiration time.Duration) AuthService {
	return &authService{
		userRepo:      userRepo,
		jwtSecret:     jwtSecret,
		jwtExpiration: jwtExpiration,
	}
}

func (auth *authService) Register(request *models.RegisterUserRequest) (*models.AuthResponse, error) {
	// Existing Check
	existing, _ := auth.userRepo.GetUserByEmail(request.Email)
	if existing != nil {
		return nil, fmt.Errorf("user with email %s already exists", request.Email)
	}

	// Password Hashing
	hashedPass, error := utils.HashPassword(request.Password)
	if error != nil {
		return nil, fmt.Errorf("failed to hash password: %v", error)
	}

	//
	user := &models.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: hashedPass,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if error = auth.userRepo.CreateUser(user); error != nil {
		return nil, fmt.Errorf("failed to create user: %v", error)
	}

	token, error := utils.GenerateJwtToken(user.ID, user.Email)
	if error != nil {
		return nil, fmt.Errorf("failed to generate token: %v", error)
	}

	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (auth *authService) Login(request *models.LoginUserRequest) (*models.AuthResponse, error) {
	user, error := auth.userRepo.GetUserByEmail(request.Email)

	if error != nil {
		return nil, fmt.Errorf("invalid creadientials")
	}

	// Password Verification
	if !utils.CheckHashedPassword(user.PasswordHash, request.Password) {
		return nil, fmt.Errorf("invalid creadientials")
	}

	// JWT Token Generation
	token, error := utils.GenerateJwtToken(user.ID, user.Email)
	if error != nil {
		return nil, fmt.Errorf("failed to generate token: %v", error)
	}

	return &models.AuthResponse{
		User:  *user,
		Token: token,
	}, nil
}

func (auth *authService) ValidateToken(tokenString string) (*models.User, error) {
	claim, error := utils.ValidateJWTToken(tokenString)

	if error != nil {
		return nil, fmt.Errorf("invalid token: %v", error)
	}

	user, error := auth.userRepo.GetUserById(claim.UserId)
	if error != nil {
		return nil, fmt.Errorf("user not found: %v", error)
	}

	return user, nil
}
