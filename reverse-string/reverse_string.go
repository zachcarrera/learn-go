package reverse

func Reverse(input string) string {
	// convert the input string into a slice of runes
	// so multibyte chars can still be reversed
	inputRunes := []rune(input)
	var reversed []rune
	for i := len(inputRunes) - 1; i >= 0; i-- {
		reversed = append(reversed, inputRunes[i])
	}
	return string(reversed)
}
