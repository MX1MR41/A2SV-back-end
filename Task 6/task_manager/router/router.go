package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default() // Create a new default Gin router

	// Define the routes for the application with the appropriate middleware for each
	// tasks
	r.GET("/tasks", middleware.Logged, controllers.GetTasks)
	r.GET("/tasks/:id", middleware.Logged, controllers.GetTaskByID)
	r.POST("/tasks", middleware.Admin, controllers.CreateTask)
	r.PUT("/tasks/:id", middleware.Admin, controllers.UpdateTask)
	r.DELETE("/tasks/:id", middleware.Admin, controllers.DeleteTask)

	// users
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", middleware.Login)
	r.GET("/users", middleware.Admin, controllers.GetUsers)
	r.POST("/users/promote/:id", middleware.Admin, controllers.Promote)

	return r
}
