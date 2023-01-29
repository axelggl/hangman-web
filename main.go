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

/*
#########################################################
# Structure pour afficher le mot sur la page à afficher #
########################################################
*/
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

/*
#######################################################
#		⬇ GERE LES DEMANDES DE L'UTILISATEUR        #
######################################################
*/
func homeHandler(w http.ResponseWriter, r *http.Request) { // gère les routes de l'application en indiquant à l'application que lorsque l'utilisateur accède à l'URL "/home", la fonction homeHandler doit être utilisée pour gérer sa demande.
	// printWOOORD(structureArray)
	a := printtheWord()
	p := Page{Valeur: a}
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
#######################################################
#		⬇ Permet d'afficher le mot sur la page       #
######################################################
*/
func printtheWord() string {

	wordArray := []string{} // on vient définir notre mot dans un tableau de string

	wordRune := []rune(word) // découpe le mot pour pouvoir l'implémenter
	for i := 0; i < len(wordRune); i++ {
		if structureArray[i].isvisible == true { // si lettre == true :
			wordArray = append(wordArray, string(wordRune[i])) // on ajoute a notre tableau de string la lettre trouvée

		} else if structureArray[i].isvisible == false { // sinon affiche "_"

			temp2 := '_'
			wordArray = append(wordArray, string(temp2)) // on ajoute un _ si la lettre est fausse

		}
	}
	/*
		#######################################################
		#	⬇ Permet de transformer le mot en une string    #
		######################################################
	*/

	result := strings.Join(wordArray, " ")
	return result // retourne le résultat en une string car le serveur veut une STRING

}
