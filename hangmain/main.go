package main

import (
	"flag"
	"fmt"
	"hangman"
	"log"
	"strings"
)

const (
	maxAttempts = 6
)

var (
	wordFile = "./../mot.txt"
	attempts int
)

func main() {
	flag.StringVar(&wordFile, "f", "mot.txt", "fichier de mots")
	flag.Parse()

	words, err := hangman.ReadWordFile(wordFile)
	if err != nil {
		log.Fatal(err)
	}
	word := hangman.Choixmot(words)
	letters := hangman.LettreauHasard(word)

	for attempts < maxAttempts {
		fmt.Println("Mot : ", strings.Join(letters, " "))
		fmt.Println("Essais restants : ", maxAttempts-attempts)

		var letter string
		fmt.Print("Choisissez une lettre : ")
		fmt.Scanln(&letter)

		if hangman.DejaEssaye(letter, letters) {
			fmt.Println("Deja essayé")
			continue
		}

		if BonneLettre(letter, word) {
			Lettres(letter, &letters, word)
		} else {
			attempts++
		}

		if hangman.Fin(letters) {
			fmt.Printf("C'est gagné :). Le mot était %s, vous avez réussi avec %d erreurs", word, attempts)
			return
		}
	}

	fmt.Println("Perdu :(. Le mot était : ", word)
}

func BonneLettre(letter string, word string) bool {
	for _, l := range word {
		if string(l) == letter {
			return true
		}
	}
	return false
}

func BonMot() {

}
func Lettres(letter string, letters *[]string, word string) {
	for i, l := range word {
		if string(l) == letter {
			(*letters)[i] = letter
		}
	}
}
