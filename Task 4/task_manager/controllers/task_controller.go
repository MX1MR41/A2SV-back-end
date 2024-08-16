package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// Define a taskService which is of type ITaskService interface
var taskService data.ITaskService

// ITaskController interface defines the methods that a TaskController type must implement
type ITaskController interface {
	GetTasks(c *gin.Context) // c is the gin context, which is used to access the request and response objects
	GetTaskByID(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

// TaskController struct defines a controller for tasks
// that will implement the ITaskController interface methods
type TaskController struct {
}

// NewTaskController creates a new TaskController and initializes the taskService
func NewTaskController() *TaskController {
	taskService = data.NewTaskService()
	return &TaskController{}
}

// GetTasks returns all tasks
func (t *TaskController) GetTasks(c *gin.Context) {
	tasks := taskService.GetTasks()
	c.JSON(http.StatusOK, tasks) // c.JSON serializes the response object into JSON and writes it to the response writer

}

// GetTaskByID returns a task by ID if it exists
func (t *TaskController) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // c.Param returns the URL parameter value by key
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"}) // gin.H is a shortcut for map[string]interface{}
		return
	}

	task, err := taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask creates a new task
func (t *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	// c.ShouldBindJSON binds the request body to the task struct
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task = taskService.CreateTask(task)
	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates a task by ID
func (t *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := taskService.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by ID
func (t *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
