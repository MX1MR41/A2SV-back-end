package Repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserCollection() *mongo.Collection {
	user_collection = client.Database("task_manager").Collection("users")
	return user_collection
}
