package routes

import (
	"todo-service/api/handler/ping"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	// Database ping
	v1.GET("/ping", ping.GET)

	// Todo routes
	v1.GET("/todos")
	v1.POST("/todos")
	v1.PATCH("/todos/:id")
	v1.DELETE("/todos/:id")
}
