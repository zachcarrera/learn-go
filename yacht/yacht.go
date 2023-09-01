package yacht

func Score(dice []int, category string) int {
	panic("Please implement the Score function")
}

func addSameNumber(dice []int, num int) int {
	var sum int
	for _, die := range dice {
		if die == num {
			sum += die
		}
	}
	return sum
}
