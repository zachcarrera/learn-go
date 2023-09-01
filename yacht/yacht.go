package yacht

func Score(dice []int, category string) int {
	var score int
	switch category {
	case "ones":
		score = addSameNumber(dice, 1)
	case "twos":
		score = addSameNumber(dice, 2)
	case "threes":
		score = addSameNumber(dice, 3)
	case "fours":
		score = addSameNumber(dice, 4)
	case "fives":
		score = addSameNumber(dice, 5)
	case "sixes":
		score = addSameNumber(dice, 6)
	case "yacht":
		if isYacht(dice) {
			score = 50
		}
	case "choice":
		for _, die := range dice {
			score += die
		}
	case "full house":
		if isFullHouse(dice) {
			for _, die := range dice {
				score += die
			}
		}
	case "big straight":
		if isBigStraight(dice) {
			score = 30
		}
	case "little straight":
		if isLittleStraight(dice) {
			score = 30
		}
	case "four of a kind":
		if sum, ok := isFourOfKind(dice); ok {
			score = sum
		}
	}
	return score
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

func isBigStraight(dice []int) bool {
	var has2, has3, has4, has5, has6 bool
	for _, die := range dice {
		switch die {
		case 2:
			has2 = true
		case 3:
			has3 = true
		case 4:
			has4 = true
		case 5:
			has5 = true
		case 6:
			has6 = true
		}
	}
	return has2 && has3 && has4 && has5 && has6
}

func isFourOfKind(dice []int) (int, bool) {
	diceCounts := make(map[int]int)
	for _, die := range dice {
		diceCounts[die]++
	}
	for die, count := range diceCounts {
		if count >= 4 {
			return die * 4, true
		}
	}
	return 0, false
}

func isFullHouse(dice []int) bool {
	diceCounts := make(map[int]int)
	for _, die := range dice {
		diceCounts[die]++
	}
	if len(diceCounts) != 2 {
		return false
	}
	for _, count := range diceCounts {
		if count != 2 && count != 3 {
			return false
		}
	}
	return true
}
