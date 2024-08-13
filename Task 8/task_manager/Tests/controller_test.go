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

type ControllerTestSuite struct {
	suite.Suite
	userRepo    *Mocks.MockUserRepository
	taskRepo    *Mocks.MockTaskRepository
	userService Usecases.IUserService
	taskService Usecases.ITaskService
	controller  controllers.IController
}

func (suite *ControllerTestSuite) SetupTest() {
	suite.userRepo = new(Mocks.MockUserRepository)
	suite.userService = Usecases.NewUserService("test_task_manager")
	suite.taskRepo = new(Mocks.MockTaskRepository)
	suite.taskService = Usecases.NewTaskService("test_task_manager")
	suite.controller = controllers.NewController()
}

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

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
