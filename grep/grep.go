package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Search(pattern string, flags, files []string) []string {
	var matches []string
	showLineNumbers, matchFiles, caseInsensitive, invertSearch, matchLine := parseFlags(flags)
	if matchLine {
		pattern = fmt.Sprintf("^%s$", pattern)
	}
	if caseInsensitive {
		pattern = fmt.Sprintf("(?i)%s", pattern)
	}
	searchExp, err := regexp.Compile(pattern)
	if err != nil {
		panic("regex failed")
	}
	comparison := buildComparison(searchExp, invertSearch)
	for _, fileName := range files {
		myFile, err := os.Open(fileName)
		if err != nil {
			panic("error opening file")
		}
		scanner := bufio.NewScanner(myFile)
		lineNumber := 1
		for scanner.Scan() {
			currentLine := scanner.Text()

			if comparison(currentLine) {
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

func buildComparison(regex *regexp.Regexp, invertSearch bool) func(string) bool {
	if invertSearch {
		return func(s string) bool { return !regex.MatchString(s) }
	}
	return regex.MatchString
}

func parseFlags(flags []string) (showLineNumbers, matchFiles, caseInsensitive, invertSearch, matchLine bool) {
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
