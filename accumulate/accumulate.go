package accumulate

// Function Accumulate maps words using the given converter function
func Accumulate(words []string, converter func(string) string) []string {
	result := make([]string, len(words))
	for i, word := range words {
		result[i] = converter(word)
	}
	return result
}
