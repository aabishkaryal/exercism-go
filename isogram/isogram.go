package isogram

import "strings"

const letters string = "abcdefghijklmnopqrstuvwxyz"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for _, letter := range letters {
		if strings.Count(word, string(letter)) > 1 {
			return false
		}
	}
	return true
}
