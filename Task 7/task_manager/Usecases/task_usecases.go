package Usecases

import (
	"errors"
	"log"
	"task_manager/Domain"
	"task_manager/Repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var task_collection = Repositories.GetTaskCollection()
var task_ctx = Repositories.GetContext()

// GetTasks retrieves all tasks from the MongoDB collection
func GetTasks() []Domain.Task {
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
func GetTaskByID(id int) (*Domain.Task, error) {
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
func CreateTask(task Domain.Task) (Domain.Task, error) {
	_, err := task_collection.InsertOne(task_ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}

// UpdateTask updates an existing task in the MongoDB collection by its ID
func UpdateTask(id int, updatedTask Domain.Task) (*Domain.Task, error) {
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
func DeleteTask(id int) error {
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
