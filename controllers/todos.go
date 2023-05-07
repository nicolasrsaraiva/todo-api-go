package controllers

import (
	"net/http"
	"text/template"

	"github.com/nicolasrsaraiva/todo-api/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosList := models.GetToDos()
	temp.ExecuteTemplate(w, "Index", todosList)
}

func NewTodo(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.FormValue("name")
	description := r.FormValue("description")
	models.CreateToDo(name, description)

	todosList := models.GetToDos()
	temp.ExecuteTemplate(w, "Index", todosList)
}
