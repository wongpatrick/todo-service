package services

import (
	"errors"
	"todo-service/api/model"
	"todo-service/api/repository"
)

func DeleteTask(taskId int) error {
	uintTaskId := uint(taskId)
	taskParams := model.TaskParams{
		Id: &uintTaskId,
	}
	// Verify if task exists
	task, err := repository.FetchTask(taskParams)
	if err != nil {
		return err
	}

	if len(task) != 0 {
		return errors.New("Cannot find task")
	}

	ptrDelete := model.Deleted
	patchTaskParams := model.Task{
		Id:     uintTaskId,
		Status: &ptrDelete,
	}
	patchErr := repository.PatchTask(patchTaskParams)
	if patchErr != nil {
		return err
	}
	return nil
}
