package reverse

func Reverse(word string) string {
	reverse := ""
	for _, c := range word {
		reverse = string(c) + reverse
	}
	return reverse
}
