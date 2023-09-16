package services

import (
	"todo-service/api/model"
	"todo-service/api/repository"
)

// Fetch Listing with params
func FetchTaskForUser(taskParams model.TaskParams) ([]model.Task, error) {
	// TODO: Authenticate User
	// Make sure user has access to only their todo list

	taskData, queryErr := repository.FetchTask(taskParams)
	if queryErr != nil {
		// LOG
		return []model.Task{}, queryErr
	}

	return taskData, nil
}
