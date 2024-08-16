package Tests

import (
	"testing"

	"task_manager/Domain"

	"github.com/stretchr/testify/assert"
)

// Test the Task struct
func TestTaskStruct(t *testing.T) {
	task := Domain.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     "2023-12-31",
		Status:      "Pending",
	}

	assert.Equal(t, 1, task.ID)
	assert.Equal(t, "Test Task", task.Title)
	assert.Equal(t, "This is a test task", task.Description)
	assert.Equal(t, "2023-12-31", task.DueDate)
	assert.Equal(t, "Pending", task.Status)
}

// Test the User struct
func TestUserStruct(t *testing.T) {
	user := Domain.User{
		ID:       1,
		Username: "testuser",
		Password: "password123",
		Role:     "admin",
	}

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "password123", user.Password)
	assert.Equal(t, "admin", user.Role)
}
