package mocks

import (
	"reflect"
	"task_API/internal/models"

	"github.com/golang/mock/gomock"
)

// MockUserRepository is mock of UserRepository
type MockUserRepository struct {
	controller *gomock.Controller
	recorder   *UserRepositoryMockRecorder
}

type UserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// New user repository mock instance form mock contorller
func NewMockUserRepository(controller *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{controller: controller}
	mock.recorder = &UserRepositoryMockRecorder{mock}
	return mock
}

// Create mocks base methods
func (mockRepo *MockUserRepository) CreateUser(user *models.User) error {
	mockRepo.controller.T.Helper()
	ret := mockRepo.controller.Call(mockRepo, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create expected call
func (mockRecorder UserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mockRecorder.mock.controller.T.Helper()
	return mockRecorder.mock.controller.RecordCallWithMethodType(mockRecorder.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// GetByEmail mock base method
func (mockRepo *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	mockRepo.controller.T.Helper()
	ret := mockRepo.controller.Call(mockRepo, "GetUserByEmail", email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByemail indicating expected call
func (mockRecorder UserRepositoryMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mockRecorder.mock.controller.T.Helper()
	return mockRecorder.mock.controller.RecordCallWithMethodType(mockRecorder.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), email)
}

// / Delete User
func (mockRepo *MockUserRepository) DeleteUser(userId int) error {
	mockRepo.controller.T.Helper()
	ret := mockRepo.controller.Call(mockRepo, "DeleteUser", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mockRecorder UserRepositoryMockRecorder) DeleteUser(userId interface{}) *gomock.Call {
	mockRecorder.mock.controller.T.Helper()
	return mockRecorder.mock.controller.RecordCallWithMethodType(mockRecorder.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), userId)
}

// Get User by id
func (mockRepo *MockUserRepository) GetUserById(id int) (*models.User, error) {
	mockRepo.controller.T.Helper()
	ret := mockRepo.controller.Call(mockRepo, "GetUserById", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mockRecorder UserRepositoryMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mockRecorder.mock.controller.T.Helper()
	return mockRecorder.mock.controller.RecordCallWithMethodType(mockRecorder.mock, "GetUserById", reflect.TypeOf((*MockUserRepository)(nil).GetUserById), id)
}

// Get User by id
func (mockRepo *MockUserRepository) UpdateUser(user *models.User) error {
	mockRepo.controller.T.Helper()
	ret := mockRepo.controller.Call(mockRepo, "UpdateUser", user)
	ret0 := ret[0].(error)
	return ret0
}

func (mockRecorder UserRepositoryMockRecorder) UpdateUser(user interface{}) *gomock.Call {
	mockRecorder.mock.controller.T.Helper()
	return mockRecorder.mock.controller.RecordCallWithMethodType(mockRecorder.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), user)
}
