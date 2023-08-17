package palindrome

import (
	"strconv"
)

// Define Product type here.
type Product struct {
	palindrome     int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	panic("Please implement the Products function")
}

func intIsPalindrome(num int) bool {
	digits := strconv.Itoa(num)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
