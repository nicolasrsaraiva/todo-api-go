package models

import (
	"github.com/nicolasrsaraiva/todo-api/database"
)

type Todo struct {
	Id          int
	Name        string
	Description string
	Done        bool
}

func CreateToDo(Name, Description string) {
	db := database.ConnectDB()
	insertTodoDB, err := db.Prepare("INSERT INTO todo(Name, Description, Done) VALUES ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	insertTodoDB.Exec(Name, Description, false)
	defer db.Close()
}

func GetToDos() []Todo {
	db := database.ConnectDB()
	selectTodoDB, err := db.Query("SELECT * FROM todo ORDER BY ID")
	if err != nil {
		panic(err.Error())
	}

	todo := Todo{}
	todosList := []Todo{}

	for selectTodoDB.Next() {
		var id int
		var name, description string
		var done bool

		err := selectTodoDB.Scan(&id, &name, &description, &done)
		if err != nil {
			panic(err.Error())
		}

		todo.Id = id
		todo.Name = name
		todo.Description = description
		todo.Done = done

		todosList = append(todosList, todo)
	}
	defer db.Close()
	return todosList
}

func UpdateToDo(newDescription string, id int) {
	db := database.ConnectDB()
	updateTodo, err := db.Prepare("UPDATE todo SET Description = $1 WHERE ID = $2")
	if err != nil {
		panic(err.Error())
	}
	updateTodo.Exec(newDescription, id)
	defer db.Close()
}

func DeleteToDo(id int) {
	db := database.ConnectDB()
	deleteTodo, err := db.Prepare("DELETE FROM todo WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}
	deleteTodo.Exec(id)
}
