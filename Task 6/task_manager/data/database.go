package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx             = context.TODO()
	task_collection *mongo.Collection
	user_collection *mongo.Collection
)

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

	task_collection = client.Database("task_manager").Collection("tasks")
	user_collection = client.Database("task_manager").Collection(("users"))
}
