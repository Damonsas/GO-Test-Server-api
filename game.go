package main

import (
	"html/template"
	"net/http"
)

// Exercice 4

type CompleteParoleData struct {
	Message1 string
	Class    string
	Message2 string
}

func CompleteParoleHandler(w http.ResponseWriter, r *http.Request) {
	data := CompleteParoleData{}
	if r.Method == http.MethodPost {
		answer1 := r.FormValue("answer1")
		answer2 := r.FormValue("answer2")
		if answer1 != "" {
			if answer1 == "want" {
				data.Message1 = "bonne réponse"
				data.Class = "success"
			} else {
				data.Message1 = "mauvaise réponse"
				data.Class = "error"
			}
		}
		if answer2 != "" {
			if answer2 == "oui qu'il n'y a" {
				data.Message2 = "bonne réponse"
				data.Class = "success"
			} else {
				data.Message2 = "mauvaise réponse"
				data.Class = "error"
			}
		}
	}
	tmpl := template.Must(template.ParseFiles("pages/game.html"))
	tmpl.Execute(w, data)
}
