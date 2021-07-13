package reverse

// Function Reverse returns the reveresed string of the given string.
func Reverse(word string) string {
	reverse := ""
	for _, c := range word {
		reverse = string(c) + reverse
	}
	return reverse
}
