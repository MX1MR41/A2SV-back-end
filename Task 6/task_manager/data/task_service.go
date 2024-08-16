package data

import (
	"errors"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	cursor, err := task_collection.Find(ctx, bson.M{})

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

func (m *TaskService) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task

	filter := bson.M{"id": id}

	err := task_collection.FindOne(ctx, filter).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *TaskService) CreateTask(task models.Task) (models.Task, error) {
	task.ID = getNextTaskID()
	_, err := task_collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}

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
	result := task_collection.FindOneAndUpdate(ctx, filter, update)
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
	filter := bson.M{"id": id}
	result, err := task_collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func getNextTaskID() int {
	var task models.Task
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := task_collection.FindOne(ctx, bson.D{}, findOptions).Decode(&task)
	if err != nil {

		return 1
	}
	return task.ID + 1
}
