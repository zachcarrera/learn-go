package lsproduct

import (
	"errors"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, errors.New("Span cannot be larger than the input length.")
	}
	if span < 0 {
		return 0, errors.New("Span must not be negative")
	}

	for _, char := range digits {
		if !unicode.IsDigit(char) {
			return 0, errors.New("Digits input must contain only digits")
		}
	}

	var maxProduct int64 = 0

	for i := 0; i < len(digits)-span+1; i++ {
		var product int64 = 1

		for j := 0; j < span; j++ {
			product *= int64([]rune(digits)[i+j] - '0')
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}
