package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// GetTasks retrieves all tasks from the MongoDB collection
func GetTasks(c *gin.Context) {
	// Call the GetTasks function from the data package to retrieve all tasks
	tasks := data.GetTasks()
	// Serialize the tasks into JSON format and return them in the response
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by its ID from the MongoDB collection
func GetTaskByID(c *gin.Context) {
	// Convert the task ID from a string to an integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	// Call the GetTaskByID function from the data package to retrieve the task by its ID
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Serialize the task into JSON format and return it in the response
	c.JSON(http.StatusOK, task)
}

// CreateTask inserts a new task into the MongoDB collection
func CreateTask(c *gin.Context) {
	var task models.Task
	// Bind the JSON data from the request body to the task struct
	if err := c.ShouldBindJSON(&task); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the CreateTask function from the data package to insert the new task
	task, _ = data.CreateTask(task)
	// Serialize the task into JSON format and return it in the response
	c.JSON(http.StatusCreated, task)
}

// UpdateTask updates an existing task in the MongoDB collection by its ID
func UpdateTask(c *gin.Context) {
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

	// Call the UpdateTask function from the data package to update the task
	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Serialize the updated task into JSON format and return it in the response
	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task from the MongoDB collection by its ID
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Return a 204 No Content response if the task was successfully deleted
	c.JSON(http.StatusNoContent, nil)
}
