package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	// build a frequency table for word and return false if
	// there is a duplicate char that is not a space or hyphen
	frequencyTable := map[rune]int{}
	for _, char := range word {
		if char == ' ' || char == '-' {
			continue
		}
		if frequencyTable[char] >= 1 {
			return false
		}
		frequencyTable[char] += 1
	}
	return true
}
