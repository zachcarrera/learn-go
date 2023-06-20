package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0.0, errors.New("invalid number")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		square, err := Square(i)
		if err != nil {
			return 0
		}
		sum += square
	}
	return sum
}
