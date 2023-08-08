package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var cipherBuilder strings.Builder
	for _, char := range strings.ToLower(s) {
		if (unicode.IsLetter(char) || unicode.IsDigit(char)) && cipherBuilder.Len()%6 == 5 {
			cipherBuilder.WriteRune(' ')
		}
		if unicode.IsLetter(char) {
			cipherBuilder.WriteRune(translateAtbashRune(char))
		} else if unicode.IsDigit(char) {
			cipherBuilder.WriteRune(char)
		}
	}
	return cipherBuilder.String()
}

func translateAtbashRune(char rune) rune {
	return 25 - char + 97 + 97
}
