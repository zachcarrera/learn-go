package piglatin

import (
	"regexp"
	"strings"
)

var (
	// regex where a string begins with a vowel or xr or yt
	beginWithVowelRe = regexp.MustCompile(`^([aeiou].|xr|yt).+$`)

	// regex where string does not begin with a vowel
	beginWithConsonantRe = regexp.MustCompile(`^([^aeiou]+)(.+)$`)

	// regex where a string contains the sequence qu after any non vowels
	containsQuRe = regexp.MustCompile(`^([^aeiou]*qu)(.+)$`)

	// regext to match when a y should be treated as a vowel
	yAsAVowelRe = regexp.MustCompile(`^([^aeiou]+|.)(y.+)$`)
)

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	translatedWords := make([]string, 0, len(words))
	for _, word := range words {
		translatedWords = append(translatedWords, Word(word))
	}
	return strings.Join(translatedWords, " ")
}

func Word(word string) string {
	if beginWithVowelRe.MatchString(word) {
		return word + "ay"
	} else if match := containsQuRe.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	} else if match := yAsAVowelRe.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	} else if match := beginWithConsonantRe.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	}
	return word
}
