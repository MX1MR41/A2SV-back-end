package controllers

import (
	"net/http"
	"strconv"
	"task_manager/Domain"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

// IController is an interface for the Controller struct
type IController interface {
	GetTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	Promote(c *gin.Context)
}

// Controller is a struct for the Controller
// Empty struct to implement the IController interface
type Controller struct{}

// NewController returns a new instance of Controller
func NewController() IController {
	return &Controller{}
}

var taskService = Usecases.NewTaskService() // taskService is an instance of TaskService
var userService = Usecases.NewUserService() // userService is an instance of UserService

// GetTasks retrieves all tasks from the MongoDB collection
func (t *Controller) GetTasks(c *gin.Context) {
	// Call the GetTasks function from the Usecases package to retrieve all tasks
	tasks := taskService.GetTasks()
	// Serialize the tasks into JSON format and return them in the response
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by its ID from the MongoDB collection
func (t *Controller) GetTaskByID(c *gin.Context) {
	// Convert the task ID from a string to an integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	// Call the GetTaskByID function from the Usecases package to retrieve the task by its ID
	task, err := taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Serialize the task into JSON format and return it in the response
	c.JSON(http.StatusOK, task)
}

// CreateTask inserts a new task into the MongoDB collection
func (t *Controller) CreateTask(c *gin.Context) {
	var task Domain.Task
	// Bind the JSON Usecases from the request body to the task struct
	if err := c.ShouldBindJSON(&task); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the CreateTask function from the Usecases package to insert the new task
	task, _ = taskService.CreateTask(task)
	// Serialize the task into JSON format and return it in the response
	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task in the MongoDB collection by its ID
func (t *Controller) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask Domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the UpdateTask function from the Usecases package to update the task
	task, err := taskService.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Serialize the updated task into JSON format and return it in the response
	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task from the MongoDB collection by its ID
func (t *Controller) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Return a 204 No Content response if the task was successfully deleted
	c.JSON(http.StatusNoContent, nil)
}

// GetUsers retrieves all users from the MongoDB collection
func (t *Controller) GetUsers(c *gin.Context) {
	users := userService.GetUsers()
	c.JSON(http.StatusOK, users)
}

// CreateUser inserts a new user into the MongoDB collection
func (t *Controller) CreateUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := userService.CreateUser(user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}

// PromoteUser promotes a user to an admin role
func (t *Controller) Promote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := userService.Promote(id); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User promoted successfully"})
}
