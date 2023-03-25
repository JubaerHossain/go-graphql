package routes

import (
	"lms/controllers"
	"net/http"
)

func Init() {

	// load all assets
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", controllers.Home)

	// load home page
}
