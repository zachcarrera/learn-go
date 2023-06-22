package reverse

func Reverse(input string) string {
	inputRunes := []rune(input)
	var reversed []rune
	for i := len(inputRunes) - 1; i >= 0; i-- {
		reversed = append(reversed, inputRunes[i])
	}
	return string(reversed)
}
