// Package raindrops implements a function to convert numbers to raindrops.
package raindrops

import "fmt"

type conversion struct {
	n           int
	replacement string
}

var conversionTable []conversion = []conversion{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert takes a number and changes it into a raindrop.
func Convert(n int) string {
	result := ""
	for _, c := range conversionTable {
		if n%c.n == 0 {
			result += c.replacement
		}
	}
	if result != "" {
		return result
	}
	return fmt.Sprint(n)
}
