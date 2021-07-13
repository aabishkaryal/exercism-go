// Package dna implements DNA type and method to generate histogram of valid nucleotides
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, c := range d {
		v, ok := h[c]
		if !ok {
			return h, errors.New("invalid nucleotide")
		}
		h[c] = v + 1
	}
	return h, nil
}
