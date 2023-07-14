package kindergarten

import (
	"errors"
	"unicode"

	"golang.org/x/exp/slices"
)

// Define the Garden type here.
type Garden struct {
	children []string
	plants   map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(children) < 1 {
		return nil, errors.New("List of children must not be empty")
	}

	isUpper := true
	for _, char := range diagram {
		if !unicode.IsUpper(char) && unicode.IsLetter(char) {
			isUpper = false
		}
	}
	if !isUpper {
		return nil, errors.New("Diagram must be formatted with capital letters")
	}

	diagramSlice := []rune(diagram)

	if len(diagramSlice) < 1 {
		return nil, errors.New("Diagram must contain atleast 1 character")
	}
	if diagramSlice[0] != '\n' {
		return nil, errors.New("Diagram must start with a new line character")
	}

	secondNewLine := slices.Index(diagramSlice[1:], '\n') + 1
	if secondNewLine == -1 {
		return nil, errors.New("Diagram must be formatted correctly")
	}

	if len(diagramSlice[1:secondNewLine])%2 != 0 || len(diagramSlice[secondNewLine+1:])%2 != 0 {
		return nil, errors.New("Diagram must have an even amount of cups in each row")
	}

	if len(diagramSlice[1:secondNewLine]) != len(diagramSlice[secondNewLine+1:]) {
		return nil, errors.New("Diagram rows must be of equal length")
	}

	return nil, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	panic("Please implement the Plants function")
}
