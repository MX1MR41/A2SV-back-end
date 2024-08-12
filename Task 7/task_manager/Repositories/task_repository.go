package Repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// ITaskRepository is an interface for TaskRepository
type ITaskRepository interface {
	GetTaskCollection() *mongo.Collection
}

// TaskRepository is a struct for TaskRepository
type TaskRepository struct{}

// NewTaskRepository returns a new instance of TaskRepository
func NewTaskRepository() ITaskRepository {
	return &TaskRepository{}
}

func (t *TaskRepository) GetTaskCollection() *mongo.Collection {
	task_collection = client.Database("task_manager").Collection("tasks")
	return task_collection
}
