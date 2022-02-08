package tasks

import (
	"github.com/gin-gonic/gin"
	"sample-todo/logic"
)

func AddRoutes(r *gin.Engine) {
	routes := r.Group("/tasks")
	routes.GET("/", logic.GetTasks)
	routes.POST("/", logic.CreateTask)
	routes.PUT("/:id", logic.UpdateTask)
	routes.DELETE("/:id", logic.DeleteTask)
}
