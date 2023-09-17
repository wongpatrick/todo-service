package task

import (
	"log"
	"net/http"
	"strconv"
	"todo-service/api/model"
	"todo-service/api/services"

	"github.com/gin-gonic/gin"
)

// TODO MODIFY THIS TO DO POST AND NOT PATCH

// POST Task
// @Summary      Create a task
// @Description  Create a task by JSON
// @Tags         task
// @Accept       json
// @Produce      json
// @Param		 task body model.Task true "Add Task"
// @Success      204  {array}   model.Task
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/task/ [post]
func POST(c *gin.Context) {
	log.Printf("PATCH")
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	var updateTask model.Task
	if err = c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	updateTask.Id = uint(aid)
	err = services.PatchTask(updateTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
