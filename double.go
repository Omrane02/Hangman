package hangman

func DejaEssaye(letter string, letters []string) bool {
	for _, l := range letters {
		if l == letter {
			return true
		}
	}
	return false
}
