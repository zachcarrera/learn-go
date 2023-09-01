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

func isYacht(dice []int) bool {
	value := dice[0]
	for _, die := range dice {
		if die != value {
			return false
		}
	}
	return true
}
