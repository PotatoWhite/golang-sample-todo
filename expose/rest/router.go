package rest

import (
	"github.com/gin-gonic/gin"
	"sample-todo/expose/rest/tasks"
)

func Init() *gin.Engine {
	r := gin.Default()
	tasks.AddRoutes(r)

	return r
}
