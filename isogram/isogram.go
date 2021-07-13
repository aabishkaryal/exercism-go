package isogram

import "strings"

const letters string = "abcdefghijklmnopqrstuvwxyz"

// Function IsIsogram returns boolean indicating if the given string is isogram.
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for _, letter := range letters {
		if strings.Count(word, string(letter)) > 1 {
			return false
		}
	}
	return true
}
