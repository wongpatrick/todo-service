package task

import (
	"log"
	"net/http"
	"todo-service/api/model"
	"todo-service/api/services"

	"github.com/gin-gonic/gin"
)

// GET Task
// @Summary      fetches task
// @Description  get task based on the user id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        q    query     model.TaskParams
// @Success      200  {array}   model.Task
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/task [get]
func GET(c *gin.Context) {
	log.Printf("FETCH")
	var taskParams model.TaskParams
	if c.ShouldBind(&taskParams) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if taskParams.Id == nil && taskParams.UserId == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	data, err := services.FetchTaskForUser(taskParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
