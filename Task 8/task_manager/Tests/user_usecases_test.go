package Tests

import (
	"errors"
	"task_manager/Domain"
	"task_manager/Tests/Mocks"
	"task_manager/Usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	userRepo    *Mocks.MockUserRepository
	userService Usecases.IUserService
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.userRepo = new(Mocks.MockUserRepository)
	suite.userService = Usecases.NewUserService("test_task_manager")

}

func (suite *UserUsecaseTestSuite) TearDownTest() {
	suite.userRepo = nil
	suite.userService = nil
}

func (suite *UserUsecaseTestSuite) TestCreateUser_ExistingUser() {

	suite.userService.CreateUser(Domain.User{
		ID:       9,
		Username: "test",
		Password: "test",
		Role:     "user",
	})

	suite.userRepo.On("CreateUser", Domain.User{

		Username: "test",
	}).Return(errors.New("user already exists"))

	err := suite.userService.CreateUser(Domain.User{Username: "test"})

	assert.NotNil(suite.T(), err)
	assert.EqualErrorf(suite.T(), err, "user already exists", "error message is not correct")

}

func (suite *UserUsecaseTestSuite) TestGetUserByUsername_UserDoesNotExist() {

	suite.userRepo.On("GetUserbyUsername", "username_that_doesnt_exist").Return(Domain.User{}, errors.New("mongo: no documents in result"))
	user, err := suite.userService.GetUserbyUsername("username_that_doesnt_exist")
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "mongo: no documents in result")
	assert.Equal(suite.T(), Domain.User{}, user)
}

func (suite *UserUsecaseTestSuite) TestPromote_UserDoesNotExist() {
	suite.userRepo.On("Promote", 99999).Return(errors.New("user not found"))
	err := suite.userService.Promote(99999)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "user not found")
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))

}
