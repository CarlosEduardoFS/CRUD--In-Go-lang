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

func Delete(w http.ResponseWriter, r *http.Request) {
	idtask := r.URL.Query().Get("id")
	services.Delete(idtask)
	http.Redirect(w, r, "/", 301)
}
