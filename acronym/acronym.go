// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	wordEnds := regexp.MustCompile(`[\s\-_]+`)
	words := wordEnds.Split(s, -1)
	var acronym []rune
	for _, v := range words {
		acronym = append(acronym, []rune(v)[0])
	}

	return string(acronym)
}
