package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	pt = strings.ToLower(pt)
	var sb strings.Builder
	for _, char := range pt {
		if unicode.IsLower(char) || unicode.IsDigit(char) {
			sb.WriteRune(char)
		}
	}
	normalized := sb.String()
	r := int(math.Sqrt(float64(len(normalized))))
	var c int
	if r*r >= len(normalized) {
		c = r
	} else {
		c = r + 1
	}

	if c*r < len(normalized) {
		r++
	}

	square := make([]string, c)
	for i := range square {
		var rowBuilder strings.Builder
		for j := i; j < len(normalized); j += c {
			rowBuilder.WriteByte(normalized[j])
		}
		square[i] = fmt.Sprintf("%-*s", r, rowBuilder.String())
	}
	return strings.Join(square, " ")
}
