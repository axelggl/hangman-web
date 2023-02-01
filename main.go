package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

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

	http.HandleFunc("/", homeRequestHandler) // utilise func homeHandler pour accéder a l'url
	// déclarer toutes tes routes : par rapport à gérer les requêtes
	fmt.Println("Localhost fonctionnel.")
	http.ListenAndServe(":8080", nil) // écoute sur le port 8080, est un gestionnaire de route
}

// fonction qui va utiliser la bonne fonction pour les méthodes http appelées par l'interface web
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

func getHome(w http.ResponseWriter, r *http.Request) { // envoie une requête a /, récupère le contenu de la page

    p := Page{Valeur: printtheWord()} // Pour le mot à trouver sur le site

    myHtml := templates.ExecuteTemplate(w, "index.html", p) // exécuter le code html
    if myHtml != nil {
        http.Error(w, myHtml.Error(), http.StatusInternalServerError) // si erreur précise erreur
    }
    
    fmt.Println(myHtml)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) { // fonction qui prévient autre type d'erreur
    w.WriteHeader(status)              // cherche le status
    if status == http.StatusNotFound { // si status est :
        fmt.Fprint(w, "custom 404") // renvoie custom 404
    }
}

func postHome(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()                   // parse le formulaire
    myLetter := r.FormValue("text") // prend la valeur du formulaire
    // le code suivant permet de modifier le mot en fonction de la lettre rentrée
    text := r.FormValue("text") // sinon écrit la page
    fmt.Println([]rune(text))

    A := []rune(myLetter) // on met input en string en rune
    for i := 0; i < len(structureArray); i++ {
        if structureArray[i].Lettre == A[0] { // regarde si input est dans le mot
            structureArray[i].isvisible = true // si c'est dans le mot isvisible devient vrai
            for o := 0; o < len(structureArray); o++ {
                if structureArray[i].Lettre == structureArray[o].Lettre { // pour les doublons
                    structureArray[o].isvisible = true
                }
            }
            fmt.Println(structureArray)
        } 
    }

    array2 := []string{}
    array := []rune(word) //tableau de structure avec mot aléatoire
    for i := 0; i < len(array); i++ {
        if structureArray[i].isvisible == true { // si condition de structure est vraie alors écrit la lettre
            array2 = append(array2, string(array[i])) // rajoute à array2 le contenu du tableau de structure
        } else if structureArray[i].isvisible == false { //sinon
            array3 := ""                         // écrit  pour cacher la lettre
            array2 = append(array2, string(array3)) // et remplace la case dans tableau 2 par _
        }
    }
    p := Page{Valeur: printtheWord()} // Pour le mot sur le site

    //renvoie le tableau de string en string grace à library join
    err := templates.ExecuteTemplate(w, "index.html", p) //exécuter le code html
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) // si erreur précise erreur
    }
}
/*
#######################################################
#		⬇ Permet d'afficher le mot sur la page       #
######################################################
*/
func printTheWord() string {

	wordArray := []string{} // on vient définir notre mot dans un tableau de string

	wordRune := []rune(word) // découpe le mot pour pouvoir l'implémenter
	for i := 0; i < len(wordRune); i++ {
		if structureArray[i].isvisible == true { // si lettre == true :
			wordArray = append(wordArray, string(wordRune[i])) // on ajoute à notre tableau de string la lettre trouvée

		} else if structureArray[i].isvisible == false { // sinon affiche "_"

			temp2 := '_'
			wordArray = append(wordArray, string(temp2)) // on ajoute un _ si la lettre est fausse

		}
	}
/*
######################################################
#	⬇ Permet de transformer le mot en une string    #
######################################################
*/

	result := strings.Join(wordArray, " ")
	return result // retourne le résultat en une string car le serveur veut une STRING

}
