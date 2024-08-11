package Repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx             = context.TODO()  // Variable that will store the context for MongoDB operations
	client          *mongo.Client     // MongoDB client
	task_collection *mongo.Collection // Variable that will hold reference to the MongoDB collection of tasks
	user_collection *mongo.Collection // Variable that will hold reference to the MongoDB collection of users
)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func GetContext() context.Context {
	return ctx
}
