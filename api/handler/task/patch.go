package task

import (
	"log"
	"net/http"
	"strconv"
	"todo-service/api/model"
	"todo-service/api/services"

	"github.com/gin-gonic/gin"
)

// PATCH Task
// @Summary      PATCH a task
// @Description  Modify a task based on the task id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id    path     int true "Task ID" Format(uint)
// @Param		 task body model.PatchTaskParams true "Update Task"
// @Success      204  {object}   model.Task
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/task/{id} [delete]
func PATCH(c *gin.Context) {
	log.Printf("PATCH")
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	var updateTask model.PatchTaskParams
	if err = c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	err = services.PatchTask(aid, updateTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
