package main

import (
	"html/template"
	"log"
	"net/http"
)

// Exercice 4

type CompleteParoleData struct {
	Message string
	Class   string
}

func CompleteParoleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("game handler called")

	data := CompleteParoleData{}
	if r.Method == http.MethodPost {
		answer1 := r.FormValue("answer1")
		answer2 := r.FormValue("answer2")
		if answer1 == "want" && answer2 == "oui qu'il n'y a" {
			data.Message = "Incroyable !"
			data.Class = "success"
		} else if answer1 == "want" || answer2 == "oui qu'il n'y a" {
			data.Message = "bonne réponse"
			data.Class = "success"
		} else {
			data.Message = "mauvaise réponse"
			data.Class = "error"
		}
	}
	tmpl := template.Must(template.ParseFiles("pages/game.html"))
	tmpl.Execute(w, data)
}
