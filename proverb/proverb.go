// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := make([]string, len(rhyme))
	for i, r := range rhyme {
		if i == len(rhyme)-1 {
			proverb[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", r, rhyme[i+1])
		}
	}
	return proverb
}
