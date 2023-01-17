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
	fmt.Println("bonjour")
	u := openWordsList()
	test := hangman_classic.CreateWord(u)
	fmt.Println(test)

	http.HandleFunc("/home", homeHandler) // utilise func homeHandler pour accèder a l'url /home
	http.ListenAndServe(":8080", nil)     // ecoute sur le port 8080,  est un gestionnaire de route
}

func homeHandler(w http.ResponseWriter, r *http.Request) { // gère les routes de l'application en indiquant à l'application que lorsque l'utilisateur accède à l'URL "/home", la fonction homeHandler doit être utilisée pour gérer sa demande.

	p := Page{Valeur: hangman_classic.CreateWord(openWordsList())}
	err := templates.ExecuteTemplate(w, "index.html", p) // prépare la réponse
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "ParseForm() err: %v", err) // gere les erreurs
	}
	Letter := r.FormValue("Letter") // attribue la variable au formulaire name
	fmt.Println(Letter)
}
