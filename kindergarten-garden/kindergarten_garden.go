package kindergarten

import (
	"errors"
	"unicode"
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

	return nil, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	panic("Please implement the Plants function")
}
