package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var sb strings.Builder
	inputChars := []rune(input)
	for i := 0; i < len(inputChars); i++ {
		currentChar := inputChars[i]
		j := i
		count := 0
		for j < len(inputChars) && inputChars[j] == currentChar {
			j++
			count++
		}
		i = j - 1
		if count > 1 {
			sb.WriteString(strconv.Itoa(count))
		}
		sb.WriteRune(currentChar)
	}
	return sb.String()
}

func RunLengthDecode(input string) string {
	var sb strings.Builder
	inputChars := []rune(input)
	for i := 0; i < len(inputChars); i++ {
		var charCount string
		for unicode.IsDigit(inputChars[i]) {
			charCount += string(inputChars[i])
			i++
		}
		if charCount == "" {
			charCount = "1"
		}
		currentChar := inputChars[i]
		convertedNum, _ := strconv.Atoi(charCount)
		for j := 0; j < convertedNum; j++ {
			sb.WriteRune(currentChar)
		}
	}
	return sb.String()
}
