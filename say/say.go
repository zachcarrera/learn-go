package say

import (
	"fmt"
	"strings"
)

var places = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
}

var numberInEnglish = map[int64]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func Say(n int64) (string, bool) {
	if n < 0 || n > 999999999999 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	numbers := separate(n)

	var words []string
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] > 0 {
			words = append(words, sayNumber(numbers[i]))
			words = append(words, places[i])
		}
	}
	return strings.TrimSpace(strings.Join(words, " ")), true
}

func separate(n int64) []int64 {
	var places []int64
	for n > 0 {
		places = append(places, n%1000)
		n /= 1000
	}
	return places
}

func sayNumber(n int64) string {
	var sb strings.Builder
	if n >= 100 {
		sb.WriteString(fmt.Sprintf("%s %s ", numberInEnglish[n/100], "hundred"))
		n %= 100
	}
	if n <= 19 && n >= 10 {
		sb.WriteString(numberInEnglish[n])
		return sb.String()
	}

	var needDash bool
	if n >= 20 {
		sb.WriteString(numberInEnglish[n/10*10])
		n %= 10
		needDash = true
	}
	if needDash && n != 0 {
		sb.WriteRune('-')
	}
	sb.WriteString(numberInEnglish[n])
	return sb.String()
}
