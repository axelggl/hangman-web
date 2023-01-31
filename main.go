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

	http.HandleFunc("/", homeRequestHandler) // utilise func homeHandler pour accèder a l'url /home
	// déclarer toute tes routes : par rapport a gérer les requetes,
	fmt.Println("Localhost fonctionnel.")
	http.ListenAndServe(":8080", nil) // ecoute sur le port 8080,  est un gestionnaire de route
}

// fonction qui vas utiliser la bonne fonction pour les méthodes http appeler par l'interface web
func homeRequestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        getHome(w, r)
        fmt.Println("get")
    } else if r.Method == "POST" {
        fmt.Println("post")
        postHome(w, r)
    } else {
        errorHandler(w, r, 404)
    }
}

func getHome(w http.ResponseWriter, r *http.Request) { // envoie une requête a /home, crée avec la struc Page et est modif avec A

    p := Page{Valeur: printtheWord()} // Pour le mot sur le site

    myHtml := templates.ExecuteTemplate(w, "index.html", p) //executer le code html
    if myHtml != nil {
        http.Error(w, myHtml.Error(), http.StatusInternalServerError) // si erreur précise erreur
    }
    
    fmt.Println(myHtml)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) { // fonction qui prévient autre type d'erreur
    w.WriteHeader(status)              // cherche le status
    if status == http.StatusNotFound { //si status est :
        fmt.Fprint(w, "custom 404") // renvoi custom 404
    }
}

func postHome(w http.ResponseWriter, r *http.Request) {

    r.ParseForm() // parse le formulaire
    myLetter := r.FormValue("text") // prend la valeur du formulaire
    //fmt.Fprint(w, "method post") // le code vient ici pour le traitement de la lettre ect.
    p := Page{Valeur: printtheWord()} // Pour le mot sur le site

    err := templates.ExecuteTemplate(w, "index.html", p) //executer le code html
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) // si erreur précise erreur
    }
    text := r.FormValue("text") // sinon écrit la page
    fmt.Println(text)
	fmt.Println(myLetter)
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
