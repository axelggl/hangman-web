package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

// Attention à utiliser précausionneusement car si 2 joueurs jouent en même temps il verront ce que font les autres.

var (
	usedLetter     []string
	word           string
	structureArray []lettre
	wordonrune     []rune
)

type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html")) // utiliser pour parser (analyse le contenu du fichier)

func main() { // gestion URL
	word = openWordsList()
	wordonrune = []rune(word)
	structureArray = cutWord()
	i := randomInt(structureArray, wordonrune)
	structureArray = i
	printWOOORD()

	http.HandleFunc("/", homeHandler) // utilise func homeHandler pour accèder a l'url /home
	// déclarer toute tes routes : par rapport a gérer les requetes,
	http.ListenAndServe(":8080", nil) // ecoute sur le port 8080,  est un gestionnaire de route
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	//Regex as needed on r.URL.Path
	//and then get the values POSTed
	Letter := r.FormValue("Letter")
	fmt.Println(Letter)

}

// #######################################################
// #		⬇ GERE LES DEMANDES DE L'UTILISATEUR         #
// ######################################################

func homeHandler(w http.ResponseWriter, r *http.Request) { // gère les routes de l'application en indiquant à l'application que lorsque l'utilisateur accède à l'URL "/home", la fonction homeHandler doit être utilisée pour gérer sa demande.
	// printWOOORD(structureArray)
	word := printWOOORD()
	p := Page{Valeur: word}
	err := templates.ExecuteTemplate(w, "index.html", p) // prépare la réponse

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "ParseForm() err: %v", err) // gere les erreurs
	}
	Letter := r.FormValue("Letter") // attribue la variable au formulaire Lettre
	// fmt.Println(Letter)
	usedLetter = append(usedLetter, Letter)
	fmt.Println(Letter)
}

/*
func getHandler(w http.ResponseWriter, r *http.Request) {

		//Match r.URL.path here as required using switch/use regex on it
	}

func postHandler(w http.ResponseWriter, r *http.Request) {

	//Regex as needed on r.URL.Path
	//and then get the values POSTed

}
//
*/
func printWOOORD() string {

	temp3 := []string{}

	temp := []rune(word)
	for i := 0; i < len(temp); i++ {
		if structureArray[i].isvisible == true { // si lettre == true :
			temp3 = append(temp3, string(temp[i]))

		} else if structureArray[i].isvisible == false { // sinon affiche "_"

			temp2 := '_'
			temp3 = append(temp3, string(temp2))

		}
	}
	sep := " "
	result := strings.Join(temp3, sep)
	return result

}
