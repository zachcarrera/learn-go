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

func isLittleStraight(dice []int) bool {
	var has1, has2, has3, has4, has5 bool
	for _, die := range dice {
		switch die {
		case 1:
			has1 = true
		case 2:
			has2 = true
		case 3:
			has3 = true
		case 4:
			has4 = true
		case 5:
			has5 = true
		}
	}
	return has1 && has2 && has3 && has4 && has5
}
