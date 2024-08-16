package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
// gin.Engine is the interface that defines the gin router
func SetupRouter() *gin.Engine {
	r := gin.Default() // Create a new gin router with default middleware

	// Create a new TaskController and assign it to the ITaskController interface
	var controller controllers.ITaskController = controllers.NewTaskController()

	// Define the routes for the application
	r.GET("/tasks", controller.GetTasks)
	r.GET("/tasks/:id", controller.GetTaskByID)
	r.POST("/tasks", controller.CreateTask)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)

	return r
}
