package protein

import "errors"

var ErrStop error = errors.New("stop codon encountered")
var ErrInvalidBase error = errors.New("invalid codon encountered")

// Function FromCodon returns the corresponding peptide for given codon
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// Function FromRNA returns the corresponding polypeptide for the given rna
func FromRNA(rna string) ([]string, error) {
	result := make([]string, 0, len(rna)/3)
	for i := 0; i < len(rna); i += 3 {
		protein := rna[i : i+3]
		codon, err := FromCodon(protein)
		if err == ErrStop {
			return result, nil
		}
		if err != nil {
			return result, err
		}
		result = append(result, codon)
	}
	return result, nil
}
