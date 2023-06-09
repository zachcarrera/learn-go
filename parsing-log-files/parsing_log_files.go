package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^(\[ERR\]|\[INF\])`)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	// ~*|\**|=*|\-*
	// re, err := regexp.Compile(`<(\-*|\**|~*|\**|\-*|=*)>`)
	re, err := regexp.Compile(`<(\-*|\**|~*|=*)>`)
	if err != nil {
		return nil
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re, err := regexp.Compile(`(?i)".*password.*"`)
	if err != nil {
		return 0
	}
	for _, line := range lines {
		if re.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {

	re, err := regexp.Compile(`end-of-line\d+`)
	if err != nil {
		return text
	}
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
