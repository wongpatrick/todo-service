package repository

import (
	"log"
	"time"
	"todo-service/api/config"
	"todo-service/api/model"

	sq "github.com/Masterminds/squirrel"
)

func PatchTask(id uint, taskParams model.PatchTaskParams) error {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return errdb
	}
	defer db.Close()

	updateQuery := buildTaskUpdate(id, taskParams)
	log.Printf(updateQuery.ToSql())
	_, err := updateQuery.RunWith(db).Query()

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func buildTaskUpdate(id uint, taskParams model.PatchTaskParams) sq.UpdateBuilder {
	update := sq.Update("task").Where("id", id)

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
