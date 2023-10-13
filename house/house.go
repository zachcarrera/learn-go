package house

import (
	"fmt"
	"strings"
)

var (
	nouns = []string{
		"the house that Jack built.",
		"the malt",
		"the rat",
		"the cat",
		"the dog",
		"the cow with the crumpled horn",
		"the maiden all forlorn",
		"the man all tattered and torn",
		"the priest all shaven and shorn",
		"the rooster that crowed in the morn",
		"the farmer sowing his corn",
		"the horse and the hound and the horn",
	}

	verbs = []string{
		"lay in",
		"ate",
		"killed",
		"worried",
		"tossed",
		"milked",
		"kissed",
		"married",
		"woke",
		"kept",
		"belonged to",
	}
)

func Verse(v int) string {
	var lines []string
	lines = append(lines, fmt.Sprintf("This is %s", nouns[v-1]))

	for i := v - 2; i >= 0; i-- {
		lines = append(lines, fmt.Sprintf("that %s %s", verbs[i], nouns[i]))
	}

	return strings.Join(lines, "\n")
}

func Song() string {
	var song []string
	for i := 1; i <= 12; i++ {
		song = append(song, Verse(i))
	}
	return strings.Join(song, "\n\n")
}
