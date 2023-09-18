package services

import (
	"errors"
	"todo-service/api/model"
	"todo-service/api/repository"
)

func DeleteTask(taskId int) error {
	taskParams := model.TaskParams{
		Id: &taskId,
	}
	// Verify if task exists
	task, err := repository.FetchTask(taskParams)
	if err != nil {
		return err
	}

	if len(task) == 0 {
		return errors.New("Cannot find task")
	}

	ptrDelete := model.Deleted
	patchTaskParams := model.PatchTaskParams{
		Status: &ptrDelete,
	}
	patchErr := repository.PatchTask(uint(taskId), patchTaskParams)
	if patchErr != nil {
		return err
	}
	return nil
}
