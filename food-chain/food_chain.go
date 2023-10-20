package foodchain

import (
	"fmt"
	"strings"
)

type verseInfo struct {
	animal  string
	comment string
	action  string
}

const prefix = "I know an old lady who swallowed a %s.\n"

var animals = []verseInfo{
	{animal: "fly"},
	{animal: "spider", comment: "It wriggled and jiggled and tickled inside her.", action: " that wriggled and jiggled and tickled inside her"},
	{animal: "bird", comment: "How absurd to swallow a bird!"},
	{animal: "cat", comment: "Imagine that, to swallow a cat!"},
	{animal: "dog", comment: "What a hog, to swallow a dog!"},
	{animal: "goat", comment: "Just opened her throat and swallowed a goat!"},
	{animal: "cow", comment: "I don't know how she swallowed a cow!"},
}

func Verse(v int) string {
	if v == 8 {
		return fmt.Sprintf(prefix, "horse") + "She's dead, of course!"
	}

	verse := fmt.Sprintf(prefix, animals[v-1].animal)
	if animals[v-1].comment != "" {
		verse += fmt.Sprintf("%s\n", animals[v-1].comment)
	}
	for i := v - 1; i > 0; i-- {
		verse += fmt.Sprintf("She swallowed the %s to catch the %s%s.\n", animals[i].animal, animals[i-1].animal, animals[i-1].action)
	}
	verse += "I don't know why she swallowed the fly. Perhaps she'll die."
	return verse
}

func Verses(start, end int) string {
	verses := make([]string, 0, end-start+1)
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, len(animals)+1)
}
