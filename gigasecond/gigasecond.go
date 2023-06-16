// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// current milliseconds
	startMilli := t.UnixMilli()

	// instead of converting to seconds and back to milliseconds
	// we can just add 1 * 10^12 (gigasecond * 1000)
	return time.UnixMilli(startMilli + 1e12)
}
