// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

package twofer

import "fmt"

// ShareWith
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	} else {
		return "One for you, one for me."
	}
}
