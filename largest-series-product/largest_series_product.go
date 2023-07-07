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

	return 0, nil
}
