package main

import (
	"html/template"
	"net/http"
)

// Exercice 4

type CompleteParoleData struct {
	Message1 string
	Class1   string
	Message2 string
	Class2   string
}

func CompleteParoleHandler(w http.ResponseWriter, r *http.Request) {
	data := CompleteParoleData{}
	if r.Method == http.MethodPost {
		answer1 := r.FormValue("answer1")
		answer2 := r.FormValue("answer2")
		if answer1 != "" {
			if answer1 == "want" {
				data.Message1 = "bonne réponse"
				data.Class1 = "success"

			} else {
				data.Message1 = "mauvaise réponse"
				data.Class1 = "error"
			}
		}
		if answer2 != "" {
			if answer2 == "oui qu'il n'y a" {
				data.Message2 = "bonne réponse"
				data.Class2 = "success"
			} else {
				data.Message2 = "mauvaise réponse"
				data.Class2 = "error"
			}
		}

	}
	tmpl := template.Must(template.ParseFiles("pages/game.html"))
	tmpl.Execute(w, data)
}
