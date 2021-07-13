// Package acronym implements functions related to acronyms
package acronym

import (
	"regexp"
	"strings"
)

// Function Abbreviate abbreviates the given string and returns it
func Abbreviate(s string) string {
	s = " " + strings.Title(s)
	acronym := ""
	re := regexp.MustCompile(`[ \-_][A-Z]`)
	for _, match := range re.FindAllString(strings.Title(s), -1) {
		acronym += string(match[1])
	}
	return acronym
}
