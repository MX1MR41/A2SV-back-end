package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

var controller = controllers.NewController()

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", Infrastructure.Logged, controller.GetTasks)
	r.GET("/tasks/:id", Infrastructure.Logged, controller.GetTaskByID)
	r.POST("/tasks", Infrastructure.Admin, controller.CreateTask)
	r.PUT("/tasks/:id", Infrastructure.Admin, controller.UpdateTask)
	r.DELETE("/tasks/:id", Infrastructure.Admin, controller.DeleteTask)

	r.POST("/register", controller.CreateUser)
	r.POST("/login", Infrastructure.Login)
	r.GET("/users", Infrastructure.Admin, controller.GetUsers)
	r.POST("/users/promote/:id", Infrastructure.Admin, controller.Promote)

	return r
}
