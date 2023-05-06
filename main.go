package main

import (
	"github.com/nicolasrsaraiva/todo-api/models"
)

func main() {
	todosList := []models.Todo{}
	models.CreateTodo("Estudar java", true, &todosList)
	models.CreateTodo("Passear com a cachorra", true, &todosList)
	models.CreateTodo("Ir para a academia", false, &todosList)
	models.CreateTodo("Almoçar bem", false, &todosList)
	models.CreateTodo("Comprar um tênis", false, &todosList)

	models.ShowTodos(todosList)
	models.ShowCompletedTodos(todosList)
	models.ShowIncompletedTodos(todosList)

}
