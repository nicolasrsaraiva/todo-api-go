package main

import (
	"github.com/nicolasrsaraiva/todo-api/models"
)

func main() {

	todosList := []models.Todo{}
	models.CreateTodo("", todosList)
}
