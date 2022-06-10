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

func Delete(id string) {
	db := db.Connect()
	delete, err := db.Prepare("delete from task where id=$1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

func Save(title string, description string, done bool) {
	db := db.Connect()

	var query = "INSERT INTO task (title,description,done) VALUES ($1,$2,$3)"

	save, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	save.Exec(title, description, done)
	defer db.Close()

}
