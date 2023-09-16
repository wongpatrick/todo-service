package routes

import (
	"todo-service/api/handler/ping"
	"todo-service/api/handler/task"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	// Database ping
	v1.GET("/ping", ping.GET)

	// Task routes
	v1.GET("/task", task.GET)
	v1.POST("/task", task.POST)
	v1.PATCH("/task/:id", task.PATCH)
	v1.DELETE("/task/:id", task.DELETE)
}
