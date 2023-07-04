package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	if !(len(isbn) == 10 || len(isbn) == 13) {
		return false
	}
	groups := strings.Split(isbn, "-")
	isbn = strings.Join(groups, "")
	sum := 0
	for i, char := range isbn {
		multiplier := 10 - i
		switch {
		case unicode.IsDigit(char):
			sum += multiplier * int(char-'0')
		case i == 9 && char == 'X':
			sum += 10
		default:
			return false
		}
	}
	return sum%11 == 0
}
