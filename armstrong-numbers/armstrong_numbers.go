package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	numberAsString := strconv.Itoa(n)
	sum := 0
	for _, num := range numberAsString {
		sum += int(math.Pow(float64(num-'0'), float64(len(numberAsString))))
	}
	return sum == n
}
