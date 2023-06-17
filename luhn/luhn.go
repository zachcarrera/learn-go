package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.Join(strings.Fields(id), "")

	if len(id) <= 1 {
		return false
	}

	var digits []int
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
		digits = append(digits, int(char-'0'))
	}

	sum := 0
	for i := range digits {
		k := len(digits) - 1 - i
		if i%2 == 1 {
			digits[k] *= 2
			if digits[k] > 9 {
				digits[k] -= 9
			}
		}
		sum += digits[k]
	}

	return sum%10 == 0
}
