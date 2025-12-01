package unit

import (
	"fmt"
	"task_API/internal/models"
	"task_API/internal/services"
	"task_API/tests/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestRegister(test *testing.T) {
// 	mockUserRepo := new(mocks.MockUserRepo)
// 	testConfig := tests.GetTestConfig()

// }

func TestService_Register(tester *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	authService := services.NewAuthService(mockUserRepo, "test_token", 24*time.Hour)

	// Register req
	registerReq := models.RegisterUserRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "Test_Password",
	}

	// Mock Calls
	mockUserRepo.On("GetUserByEmail", "test@example.com").Return(nil, fmt.Errorf("user not found"))
	mockUserRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)

	// Test Register
	authResponse, err := authService.Register(registerReq)

	// Assertions
	assert.NoError(tester, err)
	assert.NotNil(tester, authResponse)
	assert.NotEmpty(tester, authResponse.token)
	assert.Equal(tester, "test@example.com", authResponse.User.Email)

	mockUserRepo.AssertExpectations(tester)
}
