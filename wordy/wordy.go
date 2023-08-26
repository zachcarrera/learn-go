package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	if match, _ := regexp.MatchString(`What is -?\d+(?: (?:plus|minus|multiplied by|divided by) -?\d+)*\?`, question); !match {
		return 0, false
	}

	operatorsRegex := regexp.MustCompile(`(plus|minus|multiplied by|divided by)`)
	operators := operatorsRegex.FindAllString(question, -1)

	numbersRegex := regexp.MustCompile(`-?\d+`)
	numbers := numbersRegex.FindAllString(question, -1)

	sum, _ := strconv.Atoi(numbers[0])
	for i, operator := range operators {
		num, _ := strconv.Atoi(numbers[i+1])
		switch operator {
		case "plus":
			sum += num
		case "minus":
			sum -= num
		case "multiplied by":
			sum *= num
		case "divided by":
			sum /= num
		}
	}
	return sum, true
}
