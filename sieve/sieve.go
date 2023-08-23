package sieve

func Sieve(limit int) []int {
	numbersVisited := make([]bool, limit+1)
	var primes []int
	for i := 2; i <= limit; i++ {
		crossedOut := numbersVisited[i]
		if !crossedOut {
			primes = append(primes, i)
			for j := 1; ; j++ {
				k := i * j
				if k > limit {
					break
				}
				numbersVisited[k] = true
			}
		}
	}
	return primes
}
