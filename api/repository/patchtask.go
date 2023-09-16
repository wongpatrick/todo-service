package repository

import (
	"time"
	"todo-service/api/model"

	sq "github.com/Masterminds/squirrel"
)

func PatchTask(listing model.Task) error {
	// TODO: UPDATE DB with task values
	_ = buildTaskUpdate(listing)
	return nil
}

func buildTaskUpdate(taskParams model.Task) sq.UpdateBuilder {
	update := sq.Update("task").Where("id", taskParams.Id)

	update.Set("ModifiedAt", time.Now())

	if taskParams.Title != nil {
		update.Set("title", *taskParams.Title)
	}

	if taskParams.Description != nil {
		update.Set("description", *taskParams.Description)
	}

	if taskParams.Status != nil {
		update.Set("status", *taskParams.Status)
	}

	return update
}
