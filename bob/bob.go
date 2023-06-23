// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	fmt.Println("===== =====")
	fmt.Println(remark)

	// will match when strings end with a question mark and has zero or more whitespace characters
	questionRegEx := regexp.MustCompile(`^.*\?$`)
	isQuestion := questionRegEx.MatchString(remark)

	// will match when there is atleast one upper case letter and no lowercase letters
	allCapsRegEx := regexp.MustCompile(`^[A-Z1-9\s[:punct:]]+[A-Z][A-Z1-9\s[:punct:]]+$`)
	isYelling := allCapsRegEx.MatchString(remark)

	switch {
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		fmt.Printf("%s: %t\n", "question regex", questionRegEx.MatchString(remark))
		return "Sure."
	case isYelling:
		fmt.Printf("%s: %t\n", "all caps regex", allCapsRegEx.MatchString(remark))
		return "Whoa, chill out!"
	case remark == "":
		fmt.Printf("%s: %t\n", "whitespace regex", true)
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
