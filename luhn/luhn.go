package luhn

import (
	"fmt"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}
	id = strings.Join(strings.Fields(id), "")
	var digits []int
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
		digits = append(digits, int(char-'0'))
	}
	fmt.Println(digits)

	return true
}
