package main

import (
	"html/template"
	"net/http"
	"hangman-classic"
)

type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html")) // utiliser pour parser (analyse le contenu du fichier)

func main() { // gestion URL
	http.HandleFunc("/home", homeHandler) // lui dit : L'url pour envoyer la req—uette
	http.ListenAndServe(":8080", nil)     // demarre un serveur http sur le port 8080
}
func homeHandler(w http.ResponseWriter, r *http.Request) { // envoie une requête a /home, crée avec la struc Page et est modif avec A
	i := "rrbre"
	p := Page{Valeur: i}
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func postApi(w http.ResponseWriter, r *http.Request) {

}
