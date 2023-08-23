package allyourbase

import "errors"

var (
	errInputTooSmall  = errors.New("input base must be >= 2")
	errOutputTooSmall = errors.New("output base must be >= 2")
	errInvalidDigit   = errors.New("all digits must satisfy 0 <= d < input base")
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errInputTooSmall
	} else if outputBase < 2 {
		return nil, errOutputTooSmall
	}

	var inputToDecimal int
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errInvalidDigit
		}
		inputToDecimal = inputToDecimal*inputBase + digit
	}

	if inputToDecimal == 0 {
		return []int{0}, nil
	}

	var outputDigits []int
	for inputToDecimal > 0 {
		outputDigits = append([]int{inputToDecimal % outputBase}, outputDigits...)
		inputToDecimal = inputToDecimal / outputBase
	}
	return outputDigits, nil
}
