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
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	fmt.Println("===== =====")
	fmt.Println(remark)
	// will match when all charachters are not lowercase
	allCapsRegEx := regexp.MustCompile(`^[^a-z]+$`)
	if allCapsRegEx.MatchString(remark) {
		fmt.Printf("%s: %t\n", "all caps regex", allCapsRegEx.MatchString(remark))
		return "Whoa, chill out!"
	}
	return "Whatever."
}
