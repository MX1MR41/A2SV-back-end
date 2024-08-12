package data

import (
	"errors"
	"fmt"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetUsers() []models.User
	CreateUser(user models.User) error
	Promote(id int) error
	GetUserbyUsername(username string) (models.User, error)
}

type UserService struct{}

func NewUserService() IUserService {
	return &UserService{}
}

// Returns a list of all users from the database
func (m *UserService) GetUsers() []models.User {
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
func (m *UserService) CreateUser(user models.User) error {
	users := m.GetUsers()
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
	if _, err := user_collection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

// Promotes the user whose id is given into an "admin"
func (m *UserService) Promote(id int) error {
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

	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := user_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Returns use whose username matches the inputted username
func (m *UserService) GetUserbyUsername(username string) (models.User, error) {
	filter := bson.M{"username": username}
	var user models.User
	err := user_collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err

	}
	return user, nil
}

func getNextUserID() int {
	var user models.User
	findOptions := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := user_collection.FindOne(ctx, bson.D{}, findOptions).Decode(&user)
	if err != nil {
		// If no users exist, return 1 as the first ID
		return 1
	}
	return user.ID + 1
}
