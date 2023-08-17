package diamond

import (
	"errors"
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	if char > 'Z' || char < 'A' {
		return "", errors.New("Invalid char: char must be between A-Z")
	}
	sideLength := int(char-'A')*2 + 1
	matrix := make([]string, sideLength)
	for i := range matrix {
		var formattedLine string
		switch {
		case i >= 1 && i < sideLength/2+1:
			currentChar := 'A' + byte(i)
			outsideSpacing := strings.Repeat(" ", sideLength/2-i)
			insideSpacing := strings.Repeat(" ", 2*i-1)
			formattedLine = fmt.Sprintf("%s%c%s%c%s", outsideSpacing, currentChar, insideSpacing, currentChar, outsideSpacing)
		case i >= sideLength/2+1 && i < len(matrix)-1:
			index := i - sideLength/2 - 1
			currentChar := char - 1 - byte(index)
			outsideSpacing := strings.Repeat(" ", index+1)
			insideSpacing := strings.Repeat(" ", sideLength-2*index-4)
			formattedLine = fmt.Sprintf("%s%c%s%c%s", outsideSpacing, currentChar, insideSpacing, currentChar, outsideSpacing)
		default:
			spacing := strings.Repeat(" ", sideLength/2)
			formattedLine = fmt.Sprintf("%s%c%s", spacing, 'A', spacing)
		}
		matrix[i] = formattedLine
	}
	return strings.Join(matrix, "\n"), nil
}
