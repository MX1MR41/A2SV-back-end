package Usecases

import (
	"errors"
	"task_manager/Domain"
	"task_manager/Repositories"

	"golang.org/x/crypto/bcrypt"
)

var userRepo Repositories.IUserRepository

type IUserService interface {
	GetUsers() []Domain.User
	CreateUser(user Domain.User) error
	Promote(id int) error
	GetUserbyUsername(username string) (Domain.User, error)
}

type UserService struct{}

func NewUserService(dbName string) IUserService {
	userRepo = Repositories.NewUserRepository(dbName)

	return &UserService{}
}

func (u *UserService) GetUsers() []Domain.User {

	return userRepo.GetUsers()

}

func (u *UserService) CreateUser(user Domain.User) error {
	users := u.GetUsers()
	if len(users) == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	user.ID = GetNextUserID()

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
	if err := userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (u *UserService) Promote(id int) error {
	if err := userRepo.Promote(id); err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserbyUsername(username string) (Domain.User, error) {
	user, err := userRepo.GetUserbyUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetNextUserID() int {
	id := userRepo.GetNextUserID()
	return id
}
