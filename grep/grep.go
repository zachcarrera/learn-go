package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type lineComparison func(string, string) bool

func Search(pattern string, flags, files []string) []string {
	// fmt.Printf("pattern: %v\n", pattern)
	// fmt.Printf("flags: %v\n", flags)
	// fmt.Printf("files: %v\n", files)
	var matches []string
	for _, fileName := range files {
		myFile, err := os.Open(fileName)
		if err != nil {
			panic("error opening file")
		}
		scanner := bufio.NewScanner(myFile)
		// scanner.Split(bufio.ScanLines)
		lineNumber := 1
		for scanner.Scan() {
			currentLine := scanner.Text()
			if hasFlag(flags, "-i") {
				pattern = strings.ToLower(pattern)
				currentLine = strings.ToLower(currentLine)
			}

			var comparison lineComparison
			if hasFlag(flags, "-v") {
				comparison = func(s1, s2 string) bool { return !strings.Contains(s1, s2) }
			} else {
				comparison = strings.Contains
			}
			// comparison := strings.Contains

			if hasFlag(flags, "-v") && hasFlag(flags, "-x") {
				comparison = func(s, substr string) bool { return s != substr }
			} else if hasFlag(flags, "-x") {
				comparison = func(s, substr string) bool { return s == substr }
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
				// fmt.Printf("scanner.Text(): %v\n", scanner.Text())
				matches = append(matches, match)
			}
			lineNumber++
		}
		myFile.Close()
	}
	// fmt.Printf("matches: %v\n", matches)
	return matches
}

func hasFlag(flags []string, flag string) bool {
	for _, v := range flags {
		if v == flag {
			return true
		}
	}
	return false
}
