package controllers

import (
	"A2DSW/models/services"
	"html/template"
	"net/http"
	"strconv"
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

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		description := r.FormValue("description")
		done := false

		services.Save(title, description, done)
	}

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task := services.Edit(id)
	templates.ExecuteTemplate(w, "Edit", task)
}
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("description")
	done := false

	idConvert, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	services.Update(idConvert, title, description, done)
	http.Redirect(w, r, "/", 301)
}
