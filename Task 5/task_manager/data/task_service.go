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
	collection *mongo.Collection // Variable that will hold reference to the MongoDB collection
	ctx        = context.TODO()  // Variable that will store the context for MongoDB operations
)

// Initialize the MongoDB connection and set up the collection
func init() {
	// ClientOptions contains the connection URI for the MongoDB server
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, clientOptions)
	// Check if the connection was successful
	if err != nil {
		log.Fatal(err)
	}
	// Ping the MongoDB server to check if the connection is successful
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Set the collection variable to the "tasks" collection in the "task_manager" database
	// Create the collection and/or database if it does not exist
	collection = client.Database("task_manager").Collection("tasks")
}

// GetTasks retrieves all tasks from the MongoDB collection
func GetTasks() []models.Task {
	var tasks []models.Task
	// Find all tasks in the collection and store them in the cursor variable
	cursor, err := collection.Find(ctx, bson.M{})
	// Check if there was an error while finding the tasks
	if err != nil {
		log.Fatal(err)
	}
	// Iterate over all the tasks received from the collection
	for cursor.Next(ctx) {
		var task models.Task
		// De-serialize each document from the collection into a models.Task format
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
func GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	// A filter to find the document(task) that matches id
	filter := bson.M{"id": id}
	// Find the document(task) and de-serialize it to task; if errors, catch
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
func CreateTask(task models.Task) (models.Task, error) {
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}

// UpdateTask updates an existing task in the MongoDB collection by its ID
func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
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
func DeleteTask(id int) error {
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
