package services

import (
	"errors"
	"strings"
	"todo-service/api/model"
	"todo-service/api/repository"
)

// POST task
func PostTask(taskParams model.CreateTaskParams) (model.Task, error) {
	// Should verify if user can create task for this user but that is authentication
	cleanedParams, validateErr := validateTask(taskParams)
	if validateErr != nil {
		// LOG
		return model.Task{}, validateErr
	}
	createdTask, err := repository.CreateTask(cleanedParams)
	if err != nil {
		// LOG
		return model.Task{}, err
	}
	return createdTask, nil
}

func validateTask(taskParams model.CreateTaskParams) (model.CreateTaskParams, error) {
	taskParams.Description = strings.TrimSpace(taskParams.Description)
	if len(taskParams.Description) == 0 {
		return model.CreateTaskParams{}, errors.New("Empty Description")
	}
	taskParams.Title = strings.TrimSpace(taskParams.Title)
	if len(taskParams.Title) == 0 {
		return model.CreateTaskParams{}, errors.New("Empty Title")
	}
	if taskParams.UserId <= 0 {
		return model.CreateTaskParams{}, errors.New("user id cannot be 0 or lower")
	}
	return taskParams, nil
}
