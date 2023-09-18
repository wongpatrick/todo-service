package task

import (
	"log"
	"net/http"
	"todo-service/api/model"
	"todo-service/api/services"

	"github.com/gin-gonic/gin"
)

// POST Task
// @Summary      Create a task
// @Description  Create a task by JSON
// @Tags         task
// @Accept       json
// @Produce      json
// @Param		 task body model.CreateTaskParams true "Add Task"
// @Success      201  {int}  int
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/task/ [post]
func POST(c *gin.Context) {
	log.Printf("POST")
	var taskToCreate model.CreateTaskParams
	if err := c.ShouldBindJSON(&taskToCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	created, postErr := services.PostTask(taskToCreate)
	if postErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": postErr.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, created)
}
