package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	frequencyTable := map[rune]int{}
	for _, char := range input {
		frequencyTable[char] += 1
	}

	for i := int32('a'); i <= int32('z'); i++ {
		if _, ok := frequencyTable[i]; !ok {
			return false
		}
	}
	return true
}
