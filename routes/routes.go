package routes

import (
	"A2DSW/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
