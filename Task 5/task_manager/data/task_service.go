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

// Global variables for MongoDB collection and context
var (
	collection *mongo.Collection // collection represents the MongoDB collection for tasks
	ctx        = context.TODO()  // ctx is the context for the MongoDB operations
)

// Initialize the MongoDB connection and set up the collection
func init() {
	// clientOptions is a struct that contains options for the MongoDB client including the connection URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	// Connect to the MongoDB server and assign the connection to the client variable
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the collection variable to the "tasks" collection in the "task_manager" database
	collection = client.Database("task_manager").Collection("tasks")
}

type ITaskService interface {
	GetTasks() []models.Task
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	UpdateTask(id int, updatedTask models.Task) (*models.Task, error)
	DeleteTask(id int) error
}

type TaskService struct{}

func NewTaskService() ITaskService {
	return &TaskService{}
}

func (m *TaskService) GetTasks() []models.Task {
	var tasks []models.Task
	// cursor is used to iterate over the results from the Find operation
	// bson.M{} is an empty filter that matches all documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the cursor and decode each document into a Task struct
	for cursor.Next(ctx) {
		var task models.Task
		// Decode the document into the task struct
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

func (m *TaskService) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	filter := bson.M{"id": id} // filter to find the task by its ID

	// Find the task in the collection and decode it into the task struct
	err := collection.FindOne(ctx, filter).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *TaskService) CreateTask(task models.Task) (models.Task, error) {
	_, err := collection.InsertOne(ctx, task) // Insert the task document into the collection
	if err != nil {
		return task, err
	}
	return task, nil
}

func (m *TaskService) UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	filter := bson.M{"id": id} // filter to find the task by its ID
	// update is a document that sets the fields of the task to the updated values
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"dueDate":     updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	// Find the task by its ID in "filter" and update it with the new values in "update"
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

func (m *TaskService) DeleteTask(id int) error {
	filter := bson.M{"id": id} // filter to find the task by its ID
	// Delete the task that matches the filter
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
