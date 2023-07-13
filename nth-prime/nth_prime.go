package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Nth prime number cannot be calculated")
	}
	primes := []int{2}
	nextNumber := 3
	for len(primes) < n {
		isPrime := true
		for _, prime := range primes {
			if nextNumber%prime == 0 {
				isPrime = false
				nextNumber += 2
				break
			}
		}
		if isPrime {
			primes = append(primes, nextNumber)
			nextNumber += 2
		}
	}
	return primes[n-1], nil
}
