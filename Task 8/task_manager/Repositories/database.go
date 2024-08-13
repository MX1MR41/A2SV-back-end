package Repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx    = context.TODO()
	client *mongo.Client
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

func GetClient() *mongo.Client {
	return client
}
