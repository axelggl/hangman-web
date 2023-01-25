package main

import (
	"fmt"
	"net/http"
	"text/template"
	

	
)

var wordtoFind string // Attention à utiliser précausionneusement car si 2 joueurs jouent en même temps il verront ce que font les autres.
var word string
var usedLetter []string

type Page struct {
	Valeur string
}


var templates = template.Must(template.ParseFiles("index.html")) // utiliser pour parser (analyse le contenu du fichier)

func main() { // gestion URL
	hangman()
	word := openWordsList()
	fmt.Println(word)
	// cutWord(word)
	http.HandleFunc("/home", homeHandler) // utilise func homeHandler pour accèder a l'url /home
	fmt.Println(wordtoFind)
	// ici requete post
	http.ListenAndServe(":8080", nil) // ecoute sur le port 8080,  est un gestionnaire de route

}

/*
#######################################################
#		⬇ GERE LES DEMANDES DE L'UTILISATEUR         #
######################################################
*/

func homeHandler(w http.ResponseWriter, r *http.Request) { // gère les routes de l'application en indiquant à l'application que lorsque l'utilisateur accède à l'URL "/home", la fonction homeHandler doit être utilisée pour gérer sa demande.

	p := Page{Valeur: wordtoFind}
	err := templates.ExecuteTemplate(w, "index.html", p) // prépare la réponse
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "ParseForm() err: %v", err) // gere les erreurs
	}
	Letter := r.FormValue("Letter") // attribue la variable au formulaire name
	//fmt.Println(Letter)

	usedLetter = append(usedLetter, Letter)

	

}
