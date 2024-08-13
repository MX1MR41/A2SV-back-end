package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default() // Create a new default Gin router

	var controller controllers.ITaskController = controllers.NewTaskController()

	// Define the routes for the application with the appropriate middleware for each
	// tasks
	r.GET("/tasks", middleware.Logged, controller.GetTasks)
	r.GET("/tasks/:id", middleware.Logged, controller.GetTaskByID)
	r.POST("/tasks", middleware.Admin, controller.CreateTask)
	r.PUT("/tasks/:id", middleware.Admin, controller.UpdateTask)
	r.DELETE("/tasks/:id", middleware.Admin, controller.DeleteTask)

	// users
	r.POST("/register", controller.CreateUser)
	r.POST("/login", middleware.Login)
	r.GET("/users", middleware.Admin, controller.GetUsers)
	r.POST("/users/promote/:id", middleware.Admin, controller.Promote)

	return r
}
