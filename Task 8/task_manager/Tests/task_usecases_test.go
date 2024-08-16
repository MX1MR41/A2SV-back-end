package Tests

import (
	"errors"
	"fmt"
	"task_manager/Domain"
	"task_manager/Tests/Mocks"
	"task_manager/Usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and the methods that will be called in the tests
type TaskUsecaseTestSuite struct {
	suite.Suite                           // Embed the testify suite
	taskRepo    *Mocks.MockTaskRepository // Mocked repository
	taskService Usecases.ITaskService     // The service to test
}

// Setup the test suite
func (suite *TaskUsecaseTestSuite) SetupTest() {
	suite.taskRepo = new(Mocks.MockTaskRepository)                   // Create a new mock task repository
	suite.taskService = Usecases.NewTaskService("test_task_manager") // Create a new task service with mock database "test_task_manager"

}

// Tear down the test suite
func (suite *TaskUsecaseTestSuite) TearDownTest() {
	suite.taskRepo = nil
	suite.taskService = nil
}

// Since most functionalities in the task_usecases depend on GetTaskByID
// We will test GetTaskByID functionality when the task doesn't exist

func (suite *TaskUsecaseTestSuite) TestGetTaskByID_TaskDoesNotExist() {

	suite.taskRepo.On("GetTaskByID", 1).Return(Domain.Task{}, errors.New("mongo: no documents in result"))
	task, err := suite.taskService.GetTaskByID(99999)
	fmt.Println(task, err)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "mongo: no documents in result")
	assert.Equal(suite.T(), Domain.Task{}, task)
}

// Run the test suite
func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}
