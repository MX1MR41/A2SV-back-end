package Usecases

import (
	"errors"
	"fmt"
	"log"
	"task_manager/Domain"
	"task_manager/Repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var user_ctx = Repositories.GetContext()
var user_collection = Repositories.GetUserCollection()

// Returns a list of all users from the database
func GetUsers() []Domain.User {
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

// Creates user into the database
// First user will have role as "admin" by default
// Password is hashed before it is stored in the database
// Username is validated to be unique
func CreateUser(user Domain.User) error {
	users := GetUsers()
	if len(users) == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	user.ID = getNextUserID()

	user_name := user.Username

	for _, user := range users {
		if user.Username == user_name {
			return errors.New("user already exists")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	if _, err := user_collection.InsertOne(user_ctx, user); err != nil {
		return err
	}

	return nil
}

// Promotes the user whose id is given into an "admin"
func Promote(id int) error {
	filter := bson.M{"id": id}
	user := user_collection.FindOne(user_ctx, filter)

	if err := user.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			// Return the string "user not found" if the user is not found in an error format
			return fmt.Errorf("user not found")
		}
		fmt.Println("Error finding user:", err)
		return err
	}

	fmt.Println("PROMOTED USER IS ", user)
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := user_collection.UpdateOne(user_ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Returns use whose username matches the inputted username
func GetUserbyUsername(username string) (Domain.User, error) {
	filter := bson.M{"username": username}
	var user Domain.User
	err := user_collection.FindOne(user_ctx, filter).Decode(&user)
	if err != nil {
		return user, err

	}
	return user, nil
}

func getNextUserID() int {
	var user Domain.User
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := user_collection.FindOne(user_ctx, bson.D{}, findOptions).Decode(&user)
	if err != nil {
		// If no users exist, return 1 as the first ID
		return 1
	}
	return user.ID + 1
}
