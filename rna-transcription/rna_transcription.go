package strand

func ToRNA(dna string) string {
	dnaToRnaMap := map[rune]rune{
		'C': 'G',
		'G': 'C',
		'T': 'A',
		'A': 'U',
	}
	dnaCharacters := []rune(dna)

	for i, char := range dnaCharacters {
		dnaCharacters[i] = dnaToRnaMap[char]
	}

	return string(dnaCharacters)
}
