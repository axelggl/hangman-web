package main

import (
	"fmt"
	"net/http"
	"text/template"

	hangman_classic "github.com/sinjin314/hangman-classic"
)

var wordtoFind string // Attention à utiliser précausionneusement car si 2 joueurs jouent en même temps il verront ce que font les autres.
var word string
var usedLetter []string

type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html")) // utiliser pour parser (analyse le contenu du fichier)

func main() { // gestion URL

	word := openWordsList()
	fmt.Println(word)
	wordtoFind = hangman_classic.CreateWord(word)
	http.HandleFunc("/home", homeHandler) // utilise func homeHandler pour accèder a l'url /home
	fmt.Println(wordtoFind)
	// déclarer toute tes routes : par rapport a gérer les requetes, 
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
	Letter := r.FormValue("Letter") // attribue la variable au formulaire Lettre
	//fmt.Println(Letter)

	usedLetter = append(usedLetter, Letter)

	hangman_classic.IsInputOk(Letter, word, wordtoFind, &usedLetter) // & permet d'utiliser un pointeur pour avoir les lettres dans tout le programme

	fmt.Println(Letter)
	fmt.Println(word)
	fmt.Println(wordtoFind)

}
