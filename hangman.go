package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"github.com/01-edu/z01"
)

type lettre struct {
	Lettre    rune
	isvisible bool
}

func hangman() {
	bottom := 7
	bottomHangman := &bottom // récup adresse de bottom
	*bottomHangman = bottom  // pointeur bottom
	top := 0
	tophangman := &top // pareil mais pour le top
	*tophangman = top

	StrucArray := cutWord() // stock dans une variable notre word en [] de struct
	cutWord()
	word := openWordsList() // récupère un mot aleatoire

	randomInt(StrucArray, []rune(word)) // sélectionne lettre aléatoire
	for i := 10; i > 0; i-- {           // boucle pour nos 10 essais
		fmt.Printf("\n Vous avez %d essais \n\n\n", i)
		OpenHangman(top, bottom)            // affiche 1er pendu
		printWord(StrucArray, []rune(word)) // affiche le mot
		fmt.Print("Entrez une lettre : ")
		inputUser(StrucArray, []rune(word), bottomHangman, tophangman) // demande un input sur terminal

		iswin(StrucArray) // l'utilisateur, a-t-il gagné ?
		if iswin(StrucArray) == true {
			fmt.Println("\n")
			printWord(StrucArray, []rune(word)) // si oui affiche le mot
			fmt.Println("\n")
			fmt.Println("BRAVOOO vous avez trouvé le mot :", word)
			fmt.Printf("En %d essais, c'est pas mal :P", i) // Un message de bienveillance
			fmt.Println("\n")
			break
		}

	}
	if iswin(StrucArray) == false {
		fmt.Println("Dommage :/ le mot était ", word)
	}
}

// décompose le tableau en tableau de structure :
func cutWord() []lettre {
	wordArray := []rune(openWordsList())  // crée un tableau du mot
	arrayStruct := []lettre{}             // crée un tableau pour la structure
	for i := 0; i < len(wordArray); i++ { // Crée une boucle longueur du mot
		arrayLetter := lettre{ // var qui contient la struct
			Lettre:    wordArray[i], // Lettre = wordArray[i]
			isvisible: false,        // Lettre fausse donc ne s'affichera pas
		}
		arrayStruct = append(arrayStruct, arrayLetter) // découpe notre mot
	}
	return arrayStruct
}

// va choisir n case aléatoire et modifie notre struct en True
func randomInt(u []lettre, a []rune) []lettre {
	arrayStruct := u
	wordArray := a
	temp := 0
	i := rand.Intn(len(wordArray)) // Int random dans len(word)
	if len(wordArray) <= 3 {       // Si mot < 3 afficher quand même une case
		arrayStruct[i].isvisible = true
		return arrayStruct // Retourne notre tableau modifié
	}
	for {
		i := rand.Intn(len(wordArray))         // int random dans len(word)
		if arrayStruct[i].isvisible == false { // si case est fausse -> vrai pour être affiché
			arrayStruct[i].isvisible = true
			for o := 0; o < len(wordArray); o++ { // permet de vérifier si lettre en 2x
				if arrayStruct[i].Lettre == arrayStruct[o].Lettre {
					arrayStruct[o].isvisible = true
				}
			}
			temp++ // temp+1 pour éviter une loop infinity
		}
		if temp == len(wordArray)/2-1 { // si  == a notre condition, Casser la boucle, retourne le tableau
			break
		}
	}
	return arrayStruct
}

// permet d'afficher notre mot
func printWord(L []lettre, W []rune) {
	arrayStruct := L // stock notre tableau de struct
	word := W
	for i := 0; i < len(word); i++ {
		if arrayStruct[i].isvisible == true { // si lettre == true :
			z01.PrintRune(word[i]) // affiche la lettre
			z01.PrintRune(' ')
		} else if arrayStruct[i].isvisible == false { // sinon affiche "_"
			z01.PrintRune('_')
			z01.PrintRune(' ')
		}
	}
}

// demande une lettre l'utilisateur, elle la stock, et modifie notre StructArray
func inputUser(L []lettre, W []rune, bot *int, top *int) []lettre {
	Structarray := L
	word := W
	tophangman := top    // pendu = tableau
	bottomHangman := bot // permet d'afficher le pendu de "n"ligne à "n"ligne

	arr := make([]string, 0)              // crée un tableau qui va contenir la valeur l'utilisateur
	scanner := bufio.NewScanner(os.Stdin) // var qui demande l'input

	scanner.Scan()         // Ecoute le terminal :p il est gentil
	text := scanner.Text() // stock dans var l'input

	arr = append(arr, text) // lettre ajouté au tableau
	l := []rune(arr[0])     // la lettre -> rune
	temp := 0               // permet de vérifier plutard si une lettre n'est pas présente

	for i := 0; i < len(word); i++ { // cette boucle permet de vérifier si notre rune est présente
		if Structarray[i].Lettre == l[0] {
			Structarray[i].isvisible = true  // si c'est la meme alors la case devient vrai
			for o := 0; o < len(word); o++ { // pour lettre en 2x
				if Structarray[i].Lettre == Structarray[o].Lettre {
					Structarray[o].isvisible = true
				}
			}
		} else if l[0] != Structarray[i].Lettre { // permet de vérifier si lettre n'existe pas dans le mot
			temp++ // temp assure que la rune != mot.lettre
			if temp == len(word) {
				*bottomHangman = *bottomHangman + 8 // va afficher le hangman passe de n -> n+8
				*tophangman = *tophangman + 8       // passe de n  -> n+8

			}
		}
	}

	return Structarray // retourne notre tableau modifié
}

func iswin(a []lettre) bool { // permet de voir si l'utilisateur à WIN
	StrucArray := a
	isWin := true // de base il gagne
	for _, r := range StrucArray {
		if r.isvisible == false { // si une case est fausse
			isWin = false // l'utulisateur ne gagne pas :(
		}
	}
	return isWin
}

/* 1 mois pour le projet  :


     _    (^)
    (_\   |_|
     \_\  |_|
     _\_\,/_|
    (`\(_|`\|
   (`\,)  \ \
    \,)   | |     Nils Jaudon + Maerten Axel
      \__(__|



*/
