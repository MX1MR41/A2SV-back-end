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

// Define the suite, and the methods that will be called in the tests
type UserUsecaseTestSuite struct {
	suite.Suite                           // Embed the testify suite
	userRepo    *Mocks.MockUserRepository // Mocked repository
	userService Usecases.IUserService     // The service to test
}

// Setup the test suite
func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.userRepo = new(Mocks.MockUserRepository)                   // Create a new mock user repository
	suite.userService = Usecases.NewUserService("test_task_manager") // Create a new user service with mock database "test_task_manager"

}

// Tear down the test suite
func (suite *UserUsecaseTestSuite) TearDownTest() {
	suite.userRepo = nil
	suite.userService = nil
}

// Test the CreateUser method when a user with the same username already exists
func (suite *UserUsecaseTestSuite) TestCreateUser_ExistingUser() {

	suite.userService.CreateUser(Domain.User{
		ID:       9,
		Username: "test",
		Password: "test",
		Role:     "user",
	})

	// Mock the CreateUser method to return an error
	suite.userRepo.On("CreateUser", Domain.User{Username: "test"}).Return(errors.New("user already exists"))
	// Call the CreateUser method and get actual error
	err := suite.userService.CreateUser(Domain.User{Username: "test"})

	assert.NotNil(suite.T(), err)
	assert.EqualErrorf(suite.T(), err, "user already exists", "error message is not correct")

}

// Test the GetUserbyUsername method when the user does not exist
func (suite *UserUsecaseTestSuite) TestGetUserByUsername_UserDoesNotExist() {

	suite.userRepo.On("GetUserbyUsername", "username_that_doesnt_exist").Return(Domain.User{}, errors.New("mongo: no documents in result"))
	user, err := suite.userService.GetUserbyUsername("username_that_doesnt_exist")
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "mongo: no documents in result")
	assert.Equal(suite.T(), Domain.User{}, user)
}

// Test the Promote method when the user does not exist
func (suite *UserUsecaseTestSuite) TestPromote_UserDoesNotExist() {
	suite.userRepo.On("Promote", 99999).Return(errors.New("user not found"))
	err := suite.userService.Promote(99999)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "user not found")
}

// Run the test suite
func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))

}
