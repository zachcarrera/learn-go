package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for n := 1; n < limit; n++ {
		for _, v := range divisors {
			if v != 0 && n%v == 0 {
				sum += n
				break
			}
		}
	}
	return sum
}
