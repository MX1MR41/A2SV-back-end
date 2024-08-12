package Repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// IUserRepsoitory is an interface for UserRepository
type IUserRepsoitory interface {
	GetUserCollection() *mongo.Collection
}

// UserRepository is a struct for UserRepository
type UserRepository struct{}

// NewUserRepository returns a new instance of UserRepository
func NewUserRepository() IUserRepsoitory {
	return &UserRepository{}
}

func (u *UserRepository) GetUserCollection() *mongo.Collection {
	user_collection = client.Database("task_manager").Collection("users")
	return user_collection
}
