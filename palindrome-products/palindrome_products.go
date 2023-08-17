package palindrome

import (
	"errors"
	"math"
	"sort"
	"strconv"
)

// Define Product type here.
type Product struct {
	palindrome     int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	var palindromes []int
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			if intIsPalindrome(i * j) {
				palindromes = append(palindromes, i*j)
			}
		}
	}
	if len(palindromes) == 0 {
		return Product{}, Product{}, errors.New("no palindromes found")
	}
	sort.Ints(palindromes)
	return Product{
			palindromes[0],
			getFactors(fmin, fmax, palindromes[0]),
		}, Product{
			palindromes[len(palindromes)-1],
			getFactors(fmin, fmax, palindromes[len(palindromes)-1]),
		}, nil
}

func getFactors(min, max, num int) [][2]int {
	var factors [][2]int
	for i := min; float64(i) <= math.Sqrt(float64(num)); i++ {
		if num%i == 0 && num/i <= max {
			factors = append(factors, [2]int{i, num / i})
		}
	}
	return factors
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
