package controllers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.Execute(w, nil)
}
