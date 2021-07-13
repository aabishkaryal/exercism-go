// Package strands implements a function to transcript DNA to RNA.
package strand

var transcription map[rune]string = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA transcripts DNA to RNA.
func ToRNA(dna string) string {
	rna := ""
	for _, c := range dna {
		rna += transcription[c]
	}
	return rna
}
