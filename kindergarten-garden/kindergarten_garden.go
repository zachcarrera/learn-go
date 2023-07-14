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

var plantsLookup = map[rune]string{
	'V': "violets",
	'R': "radishes",
	'C': "clover",
	'G': "grass",
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

	garden := &Garden{
		children: children,
		plants:   make(map[string][]string),
	}

	for i, child := range garden.children {
		if _, ok := garden.plants[child]; ok {
			return nil, errors.New("All children must be unique")
		}
		garden.plants[child] = []string{
			plantsLookup[diagramSlice[2*i+1]],
			plantsLookup[diagramSlice[2*i+2]],
			plantsLookup[diagramSlice[secondNewLine+2*i+1]],
			plantsLookup[diagramSlice[secondNewLine+2*i+2]],
		}
	}

	return garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	panic("Please implement the Plants function")
}
