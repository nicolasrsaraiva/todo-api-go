package db

import "database/sql"

func ConnectDB() *sql.DB {
	conn := "user=postgres dbname=todo_db password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
