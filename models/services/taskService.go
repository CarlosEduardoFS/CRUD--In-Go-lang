package services

import (
	"A2DSW/db"
	models "A2DSW/models/domain"
)

func FindAll() []models.Task {
	db := db.Connect()

	taskDb, err := db.Query("select * from task")

	if err != nil {
		panic(err.Error())
	}

	task := models.Task{}
	taskAll := []models.Task{}

	for taskDb.Next() {
		var id int
		var title, description string
		var done bool

		err := taskDb.Scan(&id, &title, &description, &done)

		if err != nil {
			panic(err.Error())
		}

		task.Id = id
		task.Title = title
		task.Description = description
		task.Done = done

		taskAll = append(taskAll, task)

	}

	defer db.Close()
	return taskAll
}
