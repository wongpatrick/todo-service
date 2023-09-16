package main

// @title           To Do Web Service
// @version         1.0
// @description     This is a simple To Do web service.

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  To Do Web Service OpenAPI
// @externalDocs.url

import (
	"log"
	"time"

	"todo-service/api/routes"

	"github.com/gin-gonic/gin"
)

// A very basic Logger for the sake of the assignment
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// TODO: Initialize Database

	r := gin.Default()
	r.Use(Logger())
	routes.SetupRoutes(r)

	r.Run() // listen and serve on localhost:8080
}
