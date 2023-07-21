package kindergarten

import (
	"errors"
	"sort"
	"unicode"
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

	secondNewLine := -1
	for i, char := range diagram {
		if i == 0 && char != '\n' {
			return nil, errors.New("Diagram must start with a new line character")
		} else if i != 0 && char == '\n' {
			secondNewLine = i
		} else if !unicode.IsUpper(char) && unicode.IsLetter(char) {
			return nil, errors.New("Diagram must be formatted with capital letters")
		}
	}

	diagramSlice := []rune(diagram)

	if secondNewLine == -1 {
		return nil, errors.New("Diagram must be formatted correctly")
	}

	topRowLength := len(diagramSlice[1:secondNewLine])
	bottomRowLength := len(diagramSlice[secondNewLine+1:])

	if topRowLength%2 != 0 || bottomRowLength%2 != 0 {
		return nil, errors.New("Diagram must have an even amount of cups in each row")
	}

	if topRowLength != bottomRowLength {
		return nil, errors.New("Diagram rows must be of equal length")
	}

	orderedChildren := make([]string, len(children))
	copy(orderedChildren, children)
	sort.Strings(orderedChildren)

	garden := &Garden{
		children: orderedChildren,
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
	plants, ok := g.plants[child]
	if !ok {
		return nil, false
	}

	return plants, true
}
