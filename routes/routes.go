package routes

import (
	"net/http"

	"github.com/nicolasrsaraiva/todo-api/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new-todo", controllers.NewTodo)
	http.HandleFunc("/insert", controllers.Insert)
}
