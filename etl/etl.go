package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	scores := map[string]int{}
	for scoreNumber, letters := range in {
		for _, letter := range letters {
			letter = strings.ToLower(letter)
			scores[letter] = scoreNumber
		}
	}
	return scores
}
