package task

import (
	"log"
	"net/http"
	"strconv"
	"todo-service/api/services"

	"github.com/gin-gonic/gin"
)

// DELETE Task
// @Summary      Delete a task
// @Description  delete task based on the task id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id    path     int true "Task ID" Format(uint)
// @Success      204  {object}   model.Task
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/task/{id} [delete]
func DELETE(c *gin.Context) {
	log.Printf("DELETE")
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	err = services.DeleteTask(aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
