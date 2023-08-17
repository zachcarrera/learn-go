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
	spacing := strings.Repeat(" ", sideLength/2)
	endLines := fmt.Sprintf("%s%c%s", spacing, 'A', spacing)
	matrix[0] = endLines
	matrix[len(matrix)-1] = endLines
	if len(matrix) == 1 {
		return strings.Join(matrix, "\n"), nil
	}
	for i := range matrix[1 : sideLength/2+1] {
		i += 1
		currentChar := 'A' + byte(i)
		outsideSpacing := strings.Repeat(" ", sideLength/2-i)
		insideSpacing := strings.Repeat(" ", 2*i-1)
		matrix[i] = fmt.Sprintf("%s%c%s%c%s", outsideSpacing, currentChar, insideSpacing, currentChar, outsideSpacing)
	}

	for i := range matrix[sideLength/2+1 : len(matrix)-1] {
		currentChar := char - 1 - byte(i)
		index := i + sideLength/2 + 1
		outsideSpacing := strings.Repeat(" ", i+1)
		insideSpacing := strings.Repeat(" ", sideLength-2*i-4)
		matrix[index] = fmt.Sprintf("%s%c%s%c%s", outsideSpacing, currentChar, insideSpacing, currentChar, outsideSpacing)
	}
	return strings.Join(matrix, "\n"), nil
}
