package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	phrase = strings.TrimSpace(phrase)
	// words := strings.Split(phrase, " ")
	myRegex := regexp.MustCompile(`[\s,:!&@\$%\^"\.]`)
	words := myRegex.Split(phrase, -1)

	// var frequency map[string]int
	frequency := map[string]int{}

	for _, word := range words {
		if word != "" && word != "'" {
			frequency[word] += 1
		}
	}

	return frequency
}
