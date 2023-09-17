package services

import (
	"errors"
	"strings"
	"todo-service/api/model"
	"todo-service/api/repository"
)

// Post buyer bid after validating
func PatchTask(id int, task model.PatchTaskParams) error {
	taskParams := model.TaskParams{
		Id: &id,
	}
	// Verify if task exists
	foundTasks, err := repository.FetchTask(taskParams)
	if err != nil {
		return err
	}

	if len(foundTasks) != 0 {
		return errors.New("Cannot find task")
	}

	patchErr := repository.PatchTask(uint(id), task)
	if patchErr != nil {
		return err
	}
	return nil
}

func validatePatchTask(taskParams model.PatchTaskParams) (model.PatchTaskParams, error) {
	if taskParams.Description != nil {
		cleanDescription := strings.TrimSpace(*taskParams.Description)
		if len(cleanDescription) == 0 {
			return model.PatchTaskParams{}, errors.New("Empty Description")
		}
		taskParams.Description = &cleanDescription

	}

	if taskParams.Title != nil {
		cleanTitle := strings.TrimSpace(*taskParams.Title)
		if len(cleanTitle) == 0 {
			return model.PatchTaskParams{}, errors.New("Empty Title")
		}
		taskParams.Title = &cleanTitle
	}

	if taskParams.Status != nil && model.IsValidStatus(*taskParams.Status) {
		return model.PatchTaskParams{}, errors.New("Status is Incorrect")
	}

	return taskParams, nil
}
