package controllers

import (
	"A2DSW/models/services"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	taskAll := services.FindAll()
	templates.ExecuteTemplate(w, "Index", taskAll)
}
