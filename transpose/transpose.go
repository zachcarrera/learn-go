package transpose

type matrix [][]rune

func Transpose(input []string) []string {
	panic("Please implement the Transpose function")
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
