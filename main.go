package main

import (
	"fmt"

	"github.com/nicolasrsaraiva/todo-api/models"
)

func main() {
	models.CreateToDo("Estudar prova", true)
	for _, todo := range models.ShowToDos() {
		fmt.Println("ID:", todo.Id)
		fmt.Println("DESCRIPTION: ", todo.Description)
		fmt.Println("DONE: ", todo.Done)
		fmt.Println("----------------------------")
	}
}
