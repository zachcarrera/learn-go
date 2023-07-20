package protein

import "errors"

var ErrStop = errors.New("Stop code was matched")
var ErrInvalidBase = errors.New("Invalid Codon detected")

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	rnaChars := []rune(rna)
	for i := 0; i < len(rnaChars); i += 3 {
		codon := string(rnaChars[i : i+3])
		protein, err := FromCodon(codon)
		if err != nil {
			if err == ErrInvalidBase {
				return nil, ErrInvalidBase
			} else if err == ErrStop {
				return proteins, nil
			}
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	var protein string
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
	return protein, nil
}
