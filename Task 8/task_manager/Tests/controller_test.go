package Tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"task_manager/Delivery/controllers"
	"task_manager/Domain"
	"task_manager/Tests/Mocks"
	"task_manager/Usecases"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite
type ControllerTestSuite struct {
	suite.Suite                           // Embed the testify suite
	userRepo    *Mocks.MockUserRepository // Mocked user repository
	taskRepo    *Mocks.MockTaskRepository // Mocked task repository
	userService Usecases.IUserService     // User service
	taskService Usecases.ITaskService     // Task service
	controller  controllers.IController   // Controller
}

// Setup the test suite
func (suite *ControllerTestSuite) SetupTest() {
	suite.userRepo = new(Mocks.MockUserRepository)                   // Create a new mock user repository
	suite.userService = Usecases.NewUserService("test_task_manager") // Create a new user service with mock database "test_task_manager"
	suite.taskRepo = new(Mocks.MockTaskRepository)
	suite.taskService = Usecases.NewTaskService("test_task_manager")
	suite.controller = controllers.NewController() // Create a new controller
}

// Tear down the test suite
func (suite *ControllerTestSuite) TearDownTest() {
	suite.userRepo = nil
	suite.userService = nil
	suite.taskRepo = nil
	suite.taskService = nil
	suite.controller = nil
}

func (suite *ControllerTestSuite) TestGetTasks() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	suite.taskRepo.On("GetTasks").Return([]Domain.Task{})
	suite.controller.GetTasks(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *ControllerTestSuite) TestCreateTask() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/tasks", strings.NewReader(`{"title":"Test Task"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.taskRepo.On("CreateTask", mock.AnythingOfType("Domain.Task")).Return(Domain.Task{}, nil)
	suite.controller.CreateTask(c)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
}

func (suite *ControllerTestSuite) TestGetUsers() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	suite.userRepo.On("GetUsers").Return([]Domain.User{})
	suite.controller.GetUsers(c)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *ControllerTestSuite) TestPromote_NotAuthorized() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "999"}}

	suite.userRepo.On("Promote", 999).Return(errors.New("unauthorized"))
	suite.controller.Promote(c)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
}

func (suite *ControllerTestSuite) TestGetTaskByID_TaskDoesNotExist() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "999"}}

	suite.taskRepo.On("GetTaskByID", 999).Return(Domain.Task{}, errors.New("mongo: no documents in result"))
	suite.controller.GetTaskByID(c)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}

// Run the test suite
func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
