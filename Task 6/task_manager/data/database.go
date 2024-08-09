package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables for MongoDB collection and context
var (
	ctx             = context.TODO()  // Variable that will store the context for MongoDB operations
	task_collection *mongo.Collection // Variable that will hold reference to the MongoDB collection of tasks
	user_collection *mongo.Collection // Variable that will hold reference to the MongoDB collection of users
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
	// Set the respective collections
	// Create the collection and/or database if it does not exist
	task_collection = client.Database("task_manager").Collection("tasks")
	user_collection = client.Database("task_manager").Collection(("users"))
}
