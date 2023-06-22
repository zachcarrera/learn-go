// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"regexp"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	fmt.Println("===== =====")
	fmt.Println(remark)

	// will match for empty strings and strings with only whitespace
	whitespaceRegEx := regexp.MustCompile(`^\s*$`)
	if whitespaceRegEx.MatchString(remark) {
		fmt.Printf("%s: %t\n", "whitespace regex", whitespaceRegEx.MatchString(remark))
		return "Fine. Be that way!"
	}

	// will match when all charachters are not lowercase
	allCapsRegEx := regexp.MustCompile(`^[^a-z]+$`)
	if allCapsRegEx.MatchString(remark) {
		fmt.Printf("%s: %t\n", "all caps regex", allCapsRegEx.MatchString(remark))
		return "Whoa, chill out!"
	}
	return "Whatever."
}
