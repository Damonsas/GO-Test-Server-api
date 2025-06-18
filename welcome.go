package main

import (
	"html/template"
	"log"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome handler called")
	name := r.URL.Query().Get("name")
	tmpl := template.Must(template.ParseFiles("pages/welcome.html"))
	if name == "" {
		name = "Visiteur"
	}
	tmpl.Execute(w, name)

}

func FormValueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		if name == "" {
			name = "Visiteur"
		}
		tmpl := template.Must(template.ParseFiles("pages/welcome.html"))
		tmpl.Execute(w, name)

	}
}
