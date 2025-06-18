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
	Answer1  string
	Answer2  string
}

func CompleteParoleHandler(w http.ResponseWriter, r *http.Request) {
	data := CompleteParoleData{}
	if r.Method == http.MethodPost {
		answer1 := r.FormValue("answer1")
		answer2 := r.FormValue("answer2")
		data.Answer1 = answer1
		data.Answer2 = answer2
		if answer1 != "" {
			if answer1 == "want" || answer1 == "Want" || answer1 == "WANT" {
				data.Message1 = "bonne réponse"
				data.Class1 = "success"

			} else {
				data.Message1 = "mauvaise réponse"
				data.Class1 = "error"
			}
		}
		if answer2 != "" {
			if answer2 == "oui qu'il n'y a" || answer2 == "Oui qu'il n'y a" || answer2 == "OUI QU'IL N'Y A" {
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
