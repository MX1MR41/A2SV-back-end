package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	var controller controllers.ITaskController = controllers.NewTaskController()

	// no middleware is required for the register functionality
	r.POST("/register", controller.CreateUser)

	// login functionality is handled by middleware.login
	r.POST("/login", middleware.Login)

	// the following routes require the user to be logged in
	r.GET("/tasks", middleware.Logged, controller.GetTasks)
	r.GET("/tasks/:id", middleware.Logged, controller.GetTaskByID)

	// the following routes require the logged-in user to be an admin
	r.POST("/tasks", middleware.Admin, controller.CreateTask)
	r.PUT("/tasks/:id", middleware.Admin, controller.UpdateTask)
	r.DELETE("/tasks/:id", middleware.Admin, controller.DeleteTask)
	r.GET("/users", middleware.Admin, controller.GetUsers)
	r.POST("/users/promote/:id", middleware.Admin, controller.Promote)

	return r
}
