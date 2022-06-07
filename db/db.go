package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	con := "user=postgres dbname=tasks password=1234567 host=localhost sslmode=disable"
	db, erro := sql.Open("postgres", con)
	if erro != nil {
		panic(erro.Error())
	}
	return db
}
