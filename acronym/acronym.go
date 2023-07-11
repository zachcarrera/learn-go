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
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	s = strings.ToUpper(s)
	wordEnds := regexp.MustCompile(`[\s]`)
	words := wordEnds.Split(s, -1)
	var acronym []rune
	for _, v := range words {
		acronym = append(acronym, []rune(v)[0])
	}

	return string(acronym)
}
