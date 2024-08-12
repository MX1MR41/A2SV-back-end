package data

import (
	"context"
	"errors"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ITaskService defines the interface for task-related operations
type ITaskService interface {
	GetTasks() []models.Task
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	UpdateTask(id int, updatedTask models.Task) (*models.Task, error)
	DeleteTask(id int) error
}

// TaskService implements the ITaskService interface using MongoDB
type TaskService struct{}

func NewTaskService() ITaskService {
	return &TaskService{}
}

// Global variables for MongoDB collection and context
var (
	collection *mongo.Collection // Variable that will hold reference to the MongoDB collection
	ctx        = context.TODO()  // Variable that will store the context for MongoDB operations
)

// Initialize the MongoDB connection and set up the collection
func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("task_manager").Collection("tasks")
}

// GetTasks retrieves all tasks from the MongoDB collection
func (m *TaskService) GetTasks() []models.Task {
	var tasks []models.Task
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks
}

// GetTaskByID retrieves a task by its ID from the MongoDB collection
func (m *TaskService) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	filter := bson.M{"id": id}
	err := collection.FindOne(ctx, filter).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateTask inserts a new task into the MongoDB collection
func (m *TaskService) CreateTask(task models.Task) (models.Task, error) {
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}

// UpdateTask updates an existing task in the MongoDB collection by its ID
func (m *TaskService) UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"dueDate":     updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	var task models.Task
	if err := result.Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTask deletes a task from the MongoDB collection by its ID
func (m *TaskService) DeleteTask(id int) error {
	filter := bson.M{"id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
