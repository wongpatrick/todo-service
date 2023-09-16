package repository

import (
	"log"
	"todo-service/api/config"
	"todo-service/api/model"

	sq "github.com/Masterminds/squirrel"
)

func FetchTask(taskParams model.TaskParams) ([]model.Task, error) {
	var db, errdb = config.Connectdb()
	if errdb != nil {
		return nil, errdb
	}
	defer db.Close()

	query := buildTaskQuery(taskParams)
	log.Printf(query.ToSql())
	rows, err := query.RunWith(db).Query()

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	task := []model.Task{}

	for rows.Next() {
		var r model.Task
		err = rows.Scan(
			&r.Id,
			&r.UserId,
			&r.Title,
			&r.Description,
			&r.Status,
			&r.CreatedAt,
			&r.ModifiedAt,
		)
		if err != nil {
			log.Printf("Scan: %v", err)
			return nil, err
		}
		task = append(task, r)

	}

	return task, nil
}

func buildTaskQuery(taskParams model.TaskParams) sq.SelectBuilder {
	query := sq.Select("*").From("task")

	if taskParams.UserId != nil {
		query.Where(sq.Eq{"user_id": *taskParams.UserId})
	}

	if taskParams.Id != nil {
		query.Where(sq.Eq{"id": *taskParams.Id})
	}

	query.Where(sq.NotEq{"status": model.Deleted})

	return query
}
