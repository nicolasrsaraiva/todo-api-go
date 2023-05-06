package models

import (
	"fmt"

	"github.com/nicolasrsaraiva/todo-api/db"
)

type Todo struct {
	Id          int
	Description string
	Done        bool
}

func CreateTodo(Description string, Done bool) {
	db := db.ConnectDB()
	insertTodoDB, err := db.Prepare("insert into todo(Description, Done) values ($1, $2)")
	if err != nil {
		panic(err.Error())
	}
	insertTodoDB.Exec(Description, Done)
	defer db.Close()
}

func ShowTodos(todosList []Todo) {
	fmt.Println("All ToDos:")
	for _, todo := range todosList {
		fmt.Printf("Description: %s | Done: %t\n", todo.Description, todo.Done)
	}
}

func ShowCompletedTodos(todosList []Todo) {
	fmt.Println("All completed ToDos:")
	for _, todo := range todosList {
		if todo.Done {
			fmt.Printf("Description: %s | Done: %t\n", todo.Description, todo.Done)
		}
	}
}

func ShowIncompletedTodos(todosList []Todo) {
	fmt.Println("All incompleted ToDos:")
	for _, todo := range todosList {
		if !todo.Done {
			fmt.Printf("Description: %s | Done: %t\n", todo.Description, todo.Done)
		}
	}
}

func MakeTodoCompleted(todoCompleted *Todo, todosList *[]Todo) {
	for _, todo := range *todosList {
		if todo.Description == todoCompleted.Description {
			todo.Done = true
		}
	}
}
