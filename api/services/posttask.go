package services

import (
	"errors"
	"log"
	"strings"
	"todo-service/api/model"
	"todo-service/api/repository"
)

// POST task
func PostTask(taskParams model.CreateTaskParams) (uint, error) {
	// Should verify if user can create task for this user but that is authentication
	cleanedParams, validateErr := validatePostTask(taskParams)
	if validateErr != nil {
		// LOG
		log.Printf(validateErr.Error())
		return 0, validateErr
	}
	createdTaskId, err := repository.CreateTask(cleanedParams)
	if err != nil {
		// LOG
		log.Printf(err.Error())
		return 0, err
	}
	return createdTaskId, nil
}

func validatePostTask(taskParams model.CreateTaskParams) (model.CreateTaskParams, error) {
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
