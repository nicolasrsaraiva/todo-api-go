package main

import (
	"net/http"

	"github.com/nicolasrsaraiva/todo-api/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
