package Repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTaskCollection() *mongo.Collection {
	task_collection = client.Database("task_manager").Collection("tasks")
	return task_collection
}
