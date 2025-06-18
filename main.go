package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/game", CompleteParoleHandler)
	log.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
