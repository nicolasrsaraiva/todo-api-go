package models

import (
	"fmt"

	"github.com/nicolasrsaraiva/todo-api/database"
)

type Todo struct {
	Id          int
	Description string
	Done        bool
}

func CreateToDo(Description string, Done bool) {
	db := database.ConnectDB()
	insertTodoDB, err := db.Prepare("insert into todo(Description, Done) values ($1, $2)")
	if err != nil {
		panic(err.Error())
	}
	insertTodoDB.Exec(Description, Done)
	defer db.Close()
}

func ShowToDos() []Todo {
	db := database.ConnectDB()
	selectTodoDB, err := db.Query("select * from todo")
	if err != nil {
		panic(err.Error())
	}

	todo := Todo{}
	todosList := []Todo{}

	for selectTodoDB.Next() {
		var id int
		var description string
		var done bool

		err := selectTodoDB.Scan(&id, &description, &done)
		if err != nil {
			panic(err.Error())
		}

		todo.Id = id
		todo.Description = description
		todo.Done = done

		todosList = append(todosList, todo)
	}
	defer db.Close()
	return todosList
}

func ShowCompletedToDos(todosList []Todo) {
	fmt.Println("All completed ToDos:")
	for _, todo := range todosList {
		if todo.Done {
			fmt.Printf("Description: %s | Done: %t\n", todo.Description, todo.Done)
		}
	}
}

func ShowIncompletedToDos(todosList []Todo) {
	fmt.Println("All incompleted ToDos:")
	for _, todo := range todosList {
		if !todo.Done {
			fmt.Printf("Description: %s | Done: %t\n", todo.Description, todo.Done)
		}
	}
}

func MakeToDoCompleted(todoCompleted *Todo, todosList *[]Todo) {
	for _, todo := range *todosList {
		if todo.Description == todoCompleted.Description {
			todo.Done = true
		}
	}
}
