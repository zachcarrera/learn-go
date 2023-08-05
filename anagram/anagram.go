package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	var anagrams []string
	return anagrams
}

func isAnagram(a, b string) bool {
	if a == b || len(a) != len(b) {
		return false
	}
}
