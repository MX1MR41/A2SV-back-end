package Mocks

import (
	"task_manager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskUsecases struct {
	mock.Mock
}

func (m *MockTaskUsecases) GetTasks() []Domain.Task {
	args := m.Called()
	return args.Get(0).([]Domain.Task)
}

func (m *MockTaskUsecases) GetTaskByID(id int) (Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(Domain.Task), args.Error(1)
}

func (m *MockTaskUsecases) CreateTask(task Domain.Task) (Domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(Domain.Task), args.Error(1)
}

func (m *MockTaskUsecases) UpdateTask(id int, updatedTask Domain.Task) error {
	args := m.Called(id, updatedTask)
	return args.Error(0)
}

func (m *MockTaskUsecases) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskUsecases) GetNextTaskID() int {
	args := m.Called()
	return args.Int(0)
}
