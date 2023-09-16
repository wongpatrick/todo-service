package services

import (
	"todo-service/api/model"
	"todo-service/api/repository"
)

// POST task
func PostTask(task model.Task) (bool, error) {
	// Should verify if user can create task for this user
	err := repository.CreateTask(task)
	if err != nil {
		// LOG
		return false, err
	}
	return true, nil
}
