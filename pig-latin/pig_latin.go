package piglatin

import (
	"regexp"
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
	panic("Please implement the Sentence function")
}
