package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

// controller is an instance of Controller
var controller = controllers.NewController()

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default() // Create a new default Gin router

	// Define the routes for the application with the appropriate Infrastructure for each
	// tasks
	r.GET("/tasks", Infrastructure.Logged, controller.GetTasks)
	r.GET("/tasks/:id", Infrastructure.Logged, controller.GetTaskByID)
	r.POST("/tasks", Infrastructure.Admin, controller.CreateTask)
	r.PUT("/tasks/:id", Infrastructure.Admin, controller.UpdateTask)
	r.DELETE("/tasks/:id", Infrastructure.Admin, controller.DeleteTask)

	// users
	r.POST("/register", controller.CreateUser)
	r.POST("/login", Infrastructure.Login)
	r.GET("/users", Infrastructure.Admin, controller.GetUsers)
	r.POST("/users/promote/:id", Infrastructure.Admin, controller.Promote)

	return r
}
