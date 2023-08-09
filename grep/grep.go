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
	showLineNumbers, matchFiles, caseInsensitive, invertSearch, matchLine := findFlags(flags)
	comparison := buildComparison(invertSearch, matchLine)
	for _, fileName := range files {
		myFile, err := os.Open(fileName)
		if err != nil {
			panic("error opening file")
		}
		scanner := bufio.NewScanner(myFile)
		lineNumber := 1
		for scanner.Scan() {
			currentLine := scanner.Text()
			if caseInsensitive {
				pattern = strings.ToLower(pattern)
				currentLine = strings.ToLower(currentLine)
			}

			if comparison(currentLine, pattern) {
				if matchFiles {
					matches = append(matches, fileName)
					break
				}
				match := scanner.Text()
				if showLineNumbers {
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

func buildComparison(hasInvert, hasMatchEntireLine bool) lineComparison {
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

func findFlags(flags []string) (showLineNumbers, matchFiles, caseInsensitive, invertSearch, matchLine bool) {
	for _, flag := range flags {
		switch flag {
		case "-n":
			showLineNumbers = true
		case "-l":
			matchFiles = true
		case "-i":
			caseInsensitive = true
		case "-v":
			invertSearch = true
		case "-x":
			matchLine = true
		}
	}
	return
}

func hasFlag(flags []string, flag string) bool {
	for _, v := range flags {
		if v == flag {
			return true
		}
	}
	return false
}
