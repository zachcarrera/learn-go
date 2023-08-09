package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type lineComparison func(string, string) bool

func Search(pattern string, flags, files []string) []string {
	var matches []string
	comparison := buildComparison(flags)
	for _, fileName := range files {
		myFile, err := os.Open(fileName)
		if err != nil {
			panic("error opening file")
		}
		scanner := bufio.NewScanner(myFile)
		lineNumber := 1
		for scanner.Scan() {
			currentLine := scanner.Text()
			if hasFlag(flags, "-i") {
				pattern = strings.ToLower(pattern)
				currentLine = strings.ToLower(currentLine)
			}

			if comparison(currentLine, pattern) {
				if hasFlag(flags, "-l") {
					matches = append(matches, fileName)
					break
				}
				match := scanner.Text()
				if hasFlag(flags, "-n") {
					match = fmt.Sprintf("%d:%s", lineNumber, match)
				}
				if len(files) > 1 {
					match = fmt.Sprintf("%s:%s", fileName, match)
				}
				matches = append(matches, match)
			}
			lineNumber++
		}
		myFile.Close()
	}
	return matches
}

func buildComparison(flags []string) lineComparison {
	hasInvert := hasFlag(flags, "-v")
	hasMatchEntireLine := hasFlag(flags, "-x")
	switch {
	case hasInvert && hasMatchEntireLine:
		return func(s1, s2 string) bool { return s1 != s2 }
	case hasMatchEntireLine:
		return func(s1, s2 string) bool { return s1 == s2 }
	case hasInvert:
		return func(s1, s2 string) bool { return !strings.Contains(s1, s2) }
	default:
		return strings.Contains
	}
}

func hasFlag(flags []string, flag string) bool {
	for _, v := range flags {
		if v == flag {
			return true
		}
	}
	return false
}
