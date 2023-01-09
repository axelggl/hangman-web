package main

import (
	"fmt"
	"html/template"
	"net/http"

	hangman_classic "github.com/sinjin314/hangman-classic"
)

type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html")) // utiliser pour parser (analyse le contenu du fichier)

func main() { // gestion URL
	hangman_classic.Clear()
	word := openWordsList()
	fmt.Println(word)
	http.HandleFunc("/home", homeHandler) // lui dit : L'url pour envoyer la req—uette
	http.ListenAndServe(":8080", nil)     // demarre un serveur http sur le port 8080

}
func homeHandler(w http.ResponseWriter, r *http.Request) { // envoie une requête a /home, crée avec la struc Page et est modif avec A
	word := openWordsList()

	fmt.Println(word)
	p := Page{Valeur: hangman_classic.CreateWord(word)}

	err := templates.ExecuteTemplate(w, "index.html", p) // prépare la réponse
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func postApi(w http.ResponseWriter, r *http.Request) {

}
