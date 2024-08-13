package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	var controller controllers.ITaskController = controllers.NewTaskController()

	r.GET("/tasks", controller.GetTasks)
	r.GET("/tasks/:id", controller.GetTaskByID)
	r.POST("/tasks", controller.CreateTask)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)

	return r
}
