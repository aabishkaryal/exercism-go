package hamming

import "errors"

// Function Distance returns the hamming distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("unequal length strands")
	}
	d := 0
	s1, s2 := []rune(a), []rune(b)
	for i, _ := range s1 {
		c1, c2 := s1[i], s2[i]
		if c1 != c2 {
			d++
		}
	}
	return d, nil
}
