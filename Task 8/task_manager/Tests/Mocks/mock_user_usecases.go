package Mocks

import (
	"task_manager/Domain"

	"github.com/stretchr/testify/mock"
)

// MockUserUsecases is a mock type for the UserUsecases interface
type MockUserUsecases struct {
	mock.Mock // Embed the testify mock
}

// Define the methods that will be called in the tests
func (m *MockUserUsecases) GetUsers() []Domain.User {
	args := m.Called()
	return args.Get(0).([]Domain.User)
}

func (m *MockUserUsecases) CreateUser(user Domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUsecases) Promote(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserUsecases) GetUserbyUsername(username string) (Domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(Domain.User), args.Error(1)
}

func (m *MockUserUsecases) GetNextUserID() int {
	args := m.Called()
	return args.Int(0)
}
