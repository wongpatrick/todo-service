package repository

import (
	"log"
	"todo-service/api/config"
	"todo-service/api/model"

	sq "github.com/Masterminds/squirrel"
)

func CreateTask(task model.CreateTaskParams) (uint, error) {

	var db, errdb = config.Connectdb()
	if errdb != nil {
		return 0, errdb
	}
	defer db.Close()

	insertQuery := buildTaskInsert(task)
	log.Printf(insertQuery.ToSql())

	var id uint
	err := insertQuery.RunWith(db).QueryRow().Scan(&id)

	if err != nil {
		log.Printf(err.Error())
		return 0, err
	}

	return id, nil
}

func buildTaskInsert(taskParams model.CreateTaskParams) sq.InsertBuilder {
	insert := sq.Insert("task").
		Columns("user_id", "title", "description", "status").
		Values(taskParams.UserId, taskParams.Title, taskParams.Description, model.Active).
		Suffix("RETURNING \"id\"")

	return insert
}
