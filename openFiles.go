package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func openWordsList() string {

	//filePath := os.Args[1]
	readFile, err := os.Open("words.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	// pick de l'aléatoire :
	rand.Seed(time.Now().Unix())        // ici on check l'heure et en fonction de l'heure il prend un nombre aléatoire
	random := rand.Intn(len(fileLines)) // entre filelines[0] et fileslines[n] en fonction du temps
	a := fileLines[random]

	// quand fini pick un nombre aléatoire entre fileLines[0] & len(fileLines[-1])
	return a
}

func OpenHangman(top int, bottom int) {

	//filePath := os.Args[1]
	readFile, err := os.Open("hangman.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	//          0             n Top          ntop                nbo       top++

}
