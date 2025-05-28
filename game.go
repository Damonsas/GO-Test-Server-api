package main

import (
	"html/template"
	"log"
	"net/http"
)

// Exercice 4

type CompleteParoleData struct {
	Message string
}

func CompleteParoleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("game handler called")

	data := CompleteParoleData{}
	if r.Method == http.MethodPost {
		answer := r.FormValue("answer")
		if answer == "want" {
			data.Message = "bonne réponse"
		} else {
			data.Message = "mauvaise réponse"
		}
	}
	tmpl := template.Must(template.ParseFiles("pages/game.html"))
	tmpl.Execute(w, data)
}
