package Usecases

import (
	"errors"
	"log"
	"task_manager/Domain"
	"task_manager/Repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var task_ctx = Repositories.GetContext()

var taskRepo = Repositories.NewTaskRepository()    // taskRepo is an instance of TaskRepository
var task_collection = taskRepo.GetTaskCollection() // task_collection is a mongo collection of tasks
// ITaskService is an interface for TaskService
type ITaskService interface {
	GetTasks() []Domain.Task
	GetTaskByID(id int) (*Domain.Task, error)
	CreateTask(task Domain.Task) (Domain.Task, error)
	UpdateTask(id int, updatedTask Domain.Task) (*Domain.Task, error)
	DeleteTask(id int) error
}

// TaskService is a struct for TaskService
type TaskService struct{}

// NewTaskService returns a new instance of TaskService
func NewTaskService() ITaskService {
	return &TaskService{}
}

// GetTasks retrieves all tasks from the MongoDB collection
func (t *TaskService) GetTasks() []Domain.Task {
	var tasks []Domain.Task
	// Find all tasks in the collection and store them in the cursor variable
	cursor, err := task_collection.Find(task_ctx, bson.M{})
	// Check if there was an error while finding the tasks
	if err != nil {
		log.Fatal(err)
	}
	// Iterate over all the tasks received from the collection
	for cursor.Next(task_ctx) {
		var task Domain.Task
		// De-serialize each document from the collection into a Domain.Task format
		// and if any errors arise, log them
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
func (t *TaskService) GetTaskByID(id int) (*Domain.Task, error) {
	var task Domain.Task
	// A filter to find the document(task) that matches id
	filter := bson.M{"id": id}
	// Find the document(task) and de-serialize it to task; if errors, catch
	err := task_collection.FindOne(task_ctx, filter).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateTask inserts a new task into the MongoDB collection
func (t *TaskService) CreateTask(task Domain.Task) (Domain.Task, error) {
	task.ID = getNextTaskID()
	_, err := task_collection.InsertOne(task_ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}

// UpdateTask updates an existing task in the MongoDB collection by its ID
func (t *TaskService) UpdateTask(id int, updatedTask Domain.Task) (*Domain.Task, error) {
	filter := bson.M{"id": id}
	// This parameter will hold and update the values as inputted
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"dueDate":     updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	result := task_collection.FindOneAndUpdate(task_ctx, filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	var task Domain.Task
	if err := result.Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTask deletes a task from the MongoDB collection by its ID
func (t *TaskService) DeleteTask(id int) error {
	filter := bson.M{"id": id}
	result, err := task_collection.DeleteOne(task_ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func getNextTaskID() int {
	var task Domain.Task
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := task_collection.FindOne(task_ctx, bson.D{}, findOptions).Decode(&task)
	if err != nil {
		// If no tasks exist, return 1 as the first ID
		return 1
	}
	return task.ID + 1
}
