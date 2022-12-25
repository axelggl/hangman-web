package main

import (
//	"html/template"
//	"net/http"
"fmt"
"hangman-classic"
)

func main() {
	fmt.Println("Hello, Module!")
}
/*
type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html"))

func a() {
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":9999", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Valeur: "A"}
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postApi(w http.ResponseWriter, r *http.Request) {

}
*/