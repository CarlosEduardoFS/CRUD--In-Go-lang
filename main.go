package main

import (
	"A2DSW/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	handlerCss()
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

func handlerCss() {
	folderSystem := http.FileServer(http.Dir("./templates/"))
	pathRequest := http.StripPrefix("/templates/", folderSystem)
	http.Handle("/templates/", pathRequest)
}
