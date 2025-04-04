package routes

import (
	"task-service/controllers"

	"github.com/gin-gonic/gin"
)

// Register routes
func RegisterTaskRoutes(r *gin.Engine) {
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskByID)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}
