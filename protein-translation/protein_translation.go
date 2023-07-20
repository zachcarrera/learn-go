package protein

import "errors"

var ErrStop = errors.New("Stop code was matched")
var ErrInvalidBase = errors.New("Invalid Codon detected")

func FromRNA(rna string) ([]string, error) {
	panic("Please implement the FromRNA function")
}

func FromCodon(codon string) (string, error) {
	panic("Please implement the FromCodon function")
}
