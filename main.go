package main

import (
	"A2DSW/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
