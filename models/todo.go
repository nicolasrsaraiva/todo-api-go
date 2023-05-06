package models

import "fmt"

type Todo struct {
	Id          int
	Description string
	Done        bool
}

func CreateTodo(desc string, done bool, todosList *[]Todo) {
	todo := Todo{Description: desc, Done: done}
	*todosList = append(*todosList, todo)
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

func TodoCompleted(todoCompleted *Todo, todosList *[]Todo) {
	for _, todo := range *todosList {
		if todo.Description == todoCompleted.Description {
			todo.Done = true
		}
	}
}
