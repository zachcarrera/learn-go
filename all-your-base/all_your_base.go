package allyourbase

import "errors"

var (
	errInputTooSmall  = errors.New("input base must be >= 2")
	errOutputTooSmall = errors.New("output base must be >= 2")
	errInvalidDigit   = errors.New("all digits must satisfy 0 <= d < input base")
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	panic("Please implement the ConvertToBase function")
}
