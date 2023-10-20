package foodchain

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
	panic("Please implement the Verse function")
}

func Verses(start, end int) string {
	panic("Please implement the Verses function")
}

func Song() string {
	panic("Please implement the Song function")
}
