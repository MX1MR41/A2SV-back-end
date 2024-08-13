package Repositories

import (
	"fmt"
	"log"
	"task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user_ctx        = GetContext()
	user_collection *mongo.Collection
)

type IUserRepository interface {
	GetUsers() []Domain.User
	CreateUser(user Domain.User) error
	Promote(id int) error
	GetUserbyUsername(username string) (Domain.User, error)
	GetNextUserID() int
}

type UserRepository struct{}

func NewUserRepository(dbName string) IUserRepository {
	user_collection = client.Database(dbName).Collection("users")
	return &UserRepository{}
}

func (u *UserRepository) GetUsers() []Domain.User {
	var users []Domain.User
	cursor, err := user_collection.Find(user_ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(user_ctx) {
		var user Domain.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users
}

func (u *UserRepository) CreateUser(user Domain.User) error {
	if _, err := user_collection.InsertOne(user_ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Promote(id int) error {
	filter := bson.M{"id": id}
	user := user_collection.FindOne(user_ctx, filter)

	if err := user.Err(); err != nil {
		if err == mongo.ErrNoDocuments {

			return fmt.Errorf("user not found")
		}

		return err
	}

	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := user_collection.UpdateOne(user_ctx, filter, update)
	if err != nil {

		return err
	}
	return nil
}

func (u *UserRepository) GetUserbyUsername(username string) (Domain.User, error) {
	filter := bson.M{"username": username}
	var user Domain.User
	err := user_collection.FindOne(user_ctx, filter).Decode(&user)
	if err != nil {
		return user, err

	}
	return user, nil
}

func (u *UserRepository) GetNextUserID() int {
	var user Domain.User
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := user_collection.FindOne(user_ctx, bson.D{}, findOptions).Decode(&user)
	if err != nil {

		return 1
	}
	return user.ID + 1
}
