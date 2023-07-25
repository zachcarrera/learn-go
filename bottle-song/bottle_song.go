package bottlesong

import (
	"fmt"
	"strings"
)

var numbers = []string{
	"No",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
}

func Recite(startBottles, takeDown int) []string {
	var song []string
	verse := "%s green bottle%s hanging on the wall,"
	ending := "There'll be %s green bottle%s hanging on the wall."
	for i := startBottles; i > startBottles-takeDown; i-- {
		if i < startBottles {
			song = append(song, "")
		}
		song = append(song, []string{
			fmt.Sprintf(verse, numbers[i], pluralize(i)),
			fmt.Sprintf(verse, numbers[i], pluralize(i)),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf(ending, strings.ToLower(numbers[i-1]), pluralize(i-1)),
		}...)
	}
	return song
}

func pluralize(n int) string {
	var plural string
	if n != 1 {
		plural = "s"
	}
	return plural
}
