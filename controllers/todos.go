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
