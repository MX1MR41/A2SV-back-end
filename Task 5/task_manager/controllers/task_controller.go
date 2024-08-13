package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// Initialize the TaskService interface
var taskService data.ITaskService = data.NewTaskService()

type ITaskController interface {
	GetTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

type TaskController struct {
}

// GetTasks retrieves all tasks from the MongoDB collection
func (t *TaskController) GetTasks(c *gin.Context) {
	tasks := taskService.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by its ID from the MongoDB collection
func (t *TaskController) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask inserts a new task into the MongoDB collection
func (t *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, _ = taskService.CreateTask(task)
	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task in the MongoDB collection by its ID
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

// DeleteTask deletes a task from the MongoDB collection by its ID
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
