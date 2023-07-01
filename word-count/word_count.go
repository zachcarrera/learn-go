package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	re := regexp.MustCompile(`\w+('\w+)?`)
	words := re.FindAllString(phrase, -1)

	frequency := map[string]int{}

	for _, word := range words {
		frequency[word] += 1
	}

	return frequency
}
