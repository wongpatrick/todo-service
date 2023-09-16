package services

import (
	"errors"
	"todo-service/api/model"
	"todo-service/api/repository"
)

// Post buyer bid after validating
func PatchTask(task model.Task) error {
	taskParams := model.TaskParams{
		Id: &task.Id,
	}
	// Verify if task exists
	foundTasks, err := repository.FetchTask(taskParams)
	if err != nil {
		return err
	}

	if len(foundTasks) != 0 {
		return errors.New("Cannot find task")
	}

	patchErr := repository.PatchTask(task)
	if patchErr != nil {
		return err
	}
	return nil
}
