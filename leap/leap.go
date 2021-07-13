// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(n int) bool {
	return n%400 == 0 || (n%4 == 0 && n%100 != 0)
}
