package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates := template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Reinaldo", "reinaldo@go.dev"}
		templates.ExecuteTemplate(w, "home.html", u)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
