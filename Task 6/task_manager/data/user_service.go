package data

import (
	"errors"
	"fmt"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Returns a list of all users from the database
func GetUsers() []models.User {
	var users []models.User
	cursor, err := user_collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var user models.User
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
func CreateUser(user models.User) error {
	users := GetUsers()
	if len(users) == 0 {
		user.Role = "admin"
	}

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
	if _, err := user_collection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

// Promotes the user whose id is given into an "admin"
func Promote(id int) error {
	filter := bson.M{"id": id}
	user := user_collection.FindOne(ctx, filter)

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
	_, err := user_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Returns use whose username matches the inputted username
func GetUserbyUsername(username string) (models.User, error) {
	filter := bson.M{"username": username}
	var user models.User
	err := user_collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err

	}
	return user, nil
}
