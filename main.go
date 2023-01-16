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

	fmt.Println("bonjours")
	http.HandleFunc("/home", homeHandler) // lui dit : L'url pour envoyer la req—uette
	http.ListenAndServe(":8080", nil)     // demarre un serveur http sur le port 8080
}

func homeHandler(w http.ResponseWriter, r *http.Request) { // envoie une requête a /home, et crée avec la struc Page et est modif avec A
	word := openWordsList()
	word_hide := hangman_classic.CreateWord(word)
	p := Page{Valeur: word_hide}

	err := templates.ExecuteTemplate(w, "index.html", p) // prépare la réponse
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "ParseForm() err: %v", err) // gere les erreurs
	}
	name := r.FormValue("name") // attribue la variable au formulaire name
	fmt.Println(name)
}
