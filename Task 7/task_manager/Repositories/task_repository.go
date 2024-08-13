package Repositories

import (
	"errors"
	"log"
	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	task_ctx        = GetContext()
	task_collection *mongo.Collection
)

type ITaskRepository interface {
	GetTasks() []Domain.Task
	CreateTask(task Domain.Task) error
	GetTaskByID(id int) (Domain.Task, error)
	GetNextTaskID() int
	UpdateTask(id int, task Domain.Task) error
	DeleteTask(id int) error
}

type TaskRepository struct{}

func NewTaskRepository() ITaskRepository {
	task_collection = client.Database("task_manager").Collection("tasks")
	return &TaskRepository{}
}

func (t *TaskRepository) GetTasks() []Domain.Task {
	var tasks []Domain.Task
	cursor, err := task_collection.Find(task_ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(task_ctx) {
		var task Domain.Task
		if err := cursor.Decode(&task); err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func (t *TaskRepository) CreateTask(task Domain.Task) error {
	if _, err := task_collection.InsertOne(task_ctx, task); err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) GetTaskByID(id int) (Domain.Task, error) {
	filter := bson.M{"id": id}
	var task Domain.Task
	if err := task_collection.FindOne(task_ctx, filter).Decode(&task); err != nil {
		return task, err
	}

	return task, nil
}

func (t *TaskRepository) GetNextTaskID() int {
	var task Domain.Task
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := task_collection.FindOne(task_ctx, bson.D{}, findOptions).Decode(&task)
	if err != nil {

		return 1
	}
	return task.ID + 1
}

func (t *TaskRepository) UpdateTask(id int, task Domain.Task) error {
	filter := bson.M{"id": id}

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"dueDate":     task.DueDate,
			"status":      task.Status,
		},
	}
	result := task_collection.FindOneAndUpdate(task_ctx, filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		return errors.New("task not found")
	}
	if result.Err() != nil {
		return result.Err()
	}
	var updatedTask Domain.Task
	if err := result.Decode(&updatedTask); err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) DeleteTask(id int) error {
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
