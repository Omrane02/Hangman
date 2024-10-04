package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

var words []string

func ReadWordFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func Choixmot(words []string) string {
	rand.Seed(time.Now().UnixNano())
	if len(words) == 0 {
		return ""
	}
	len_words := len(words)

	return words[rand.Intn(len_words)]
}

func LettreauHasard(word string) []string {
	arrWord := strings.SplitAfter(word, "")
	for index, _ := range arrWord {
		if index == random1 || index == random2 {
			continue
		}
		arrWord[index] = "_"
	}
	return arrWord
}

var random1 int
var random2 int
