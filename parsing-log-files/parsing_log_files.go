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
	return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
