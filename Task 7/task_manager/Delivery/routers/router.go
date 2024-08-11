package routers

import (
	"task_manager/Delivery/controllers"
	Infrastucture "task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default() // Create a new default Gin router

	// Define the routes for the application with the appropriate Infrastructure for each
	// tasks
	r.GET("/tasks", Infrastucture.Logged, controllers.GetTasks)
	r.GET("/tasks/:id", Infrastucture.Logged, controllers.GetTaskByID)
	r.POST("/tasks", Infrastucture.Admin, controllers.CreateTask)
	r.PUT("/tasks/:id", Infrastucture.Admin, controllers.UpdateTask)
	r.DELETE("/tasks/:id", Infrastucture.Admin, controllers.DeleteTask)

	// users
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", Infrastucture.Login)
	r.GET("/users", Infrastucture.Admin, controllers.GetUsers)
	r.POST("/users/promote/:id", Infrastucture.Admin, controllers.Promote)

	return r
}
