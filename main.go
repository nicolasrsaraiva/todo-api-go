package main

import (
	"fmt"

	"github.com/nicolasrsaraiva/todo-api/models"
)

func main() {
	models.CreateToDo("Estudar prova", true)
	models.UpdateToDo("Teste update todo", 1)
	models.DeleteToDo(3)

	for _, todo := range models.GetToDos() {
		fmt.Println("ID:", todo.Id)
		fmt.Println("DESCRIPTION: ", todo.Description)
		fmt.Println("DONE: ", todo.Done)
		fmt.Println("----------------------------")
	}
}
