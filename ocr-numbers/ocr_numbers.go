package ocr

import "strings"

var digits = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func Recognize(string) []string {
	panic("Please implement the Recognize function")
}

func recognizeLine(lines []string) string {
	var sb strings.Builder
	for i := 0; i < len(lines[0]); i += 3 {
		sb.WriteString(recognizeDigit(lines[0][i:i+3] + lines[1][i:i+3] + lines[2][i:i+3]))
	}
	return sb.String()
}

func recognizeDigit(digit string) string {
	if num, ok := digits[digit]; ok {
		return num
	}
	return "?"
}
