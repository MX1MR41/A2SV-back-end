package controllers

import (
	"net/http"
	"strconv"
	"task_manager/Domain"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

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

type Controller struct{}

func NewController() IController {
	return &Controller{}
}

var taskService Usecases.ITaskService = Usecases.NewTaskService()
var userService Usecases.IUserService = Usecases.NewUserService()

func (t *Controller) GetTasks(c *gin.Context) {

	tasks := taskService.GetTasks()

	c.JSON(http.StatusOK, tasks)
}

func (t *Controller) GetTaskByID(c *gin.Context) {

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

func (t *Controller) CreateTask(c *gin.Context) {
	var task Domain.Task

	if err := c.ShouldBindJSON(&task); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, _ = taskService.CreateTask(task)

	c.JSON(http.StatusCreated, task)
}

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

	if err := taskService.UpdateTask(id, updatedTask); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

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

	c.JSON(http.StatusNoContent, nil)
}

func (t *Controller) GetUsers(c *gin.Context) {
	users := userService.GetUsers()
	c.JSON(http.StatusOK, users)
}

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
