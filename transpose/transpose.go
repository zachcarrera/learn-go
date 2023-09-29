package transpose

type matrix [][]rune

func Transpose(input []string) []string {
	var transposed matrix
	for r, row := range input {
		for c, char := range row {
			if len(transposed) <= c {
				transposed = append(transposed, []rune{})
			}
			if len(transposed[c]) < r {
				transposed[c] = append(transposed[c], setPadding(r-len(transposed[c]), ' ')...)
			}
			transposed[c] = append(transposed[c], char)
		}
	}
	return toStringSlice(transposed)
}

func setPadding(size int, fill rune) []rune {
	var padding []rune
	for i := 0; i < size; i++ {
		padding = append(padding, fill)
	}
	return padding
}

func toStringSlice(input [][]rune) []string {
	var output []string
	for _, row := range input {
		output = append(output, string(row))
	}
	return output
}
