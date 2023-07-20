package prime

func Factors(n int64) []int64 {
	var primeFactors []int64
	factor := int64(2)
	for factor <= n {
		if n%factor == 0 {
			n /= factor
			primeFactors = append(primeFactors, factor)
		} else {
			factor++
		}
	}
	return primeFactors
}
