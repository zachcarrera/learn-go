package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	var anagrams []string
	for _, candidate := range candidates {
		if isAnagram(subject, strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func isAnagram(a, b string) bool {
	if a == b || len(a) != len(b) {
		return false
	}
	for _, char := range a {
		if strings.Count(a, string(char)) != strings.Count(b, string(char)) {
			return false
		}
	}
	return true
}
