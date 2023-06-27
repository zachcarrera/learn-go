package twelve

import (
	"fmt"
	"strings"
)

var (
	gifts = []string{
		"a Partridge in a Pear Tree",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}

	days = []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}
)

func Verse(i int) string {
	prefix := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i-1])
	if i == 1 {
		return fmt.Sprintf("%s %s.", prefix, gifts[i-1])
	}
	var list string
	for index := i; index >= 1; index-- {
		if index == 1 {
			list += fmt.Sprintf("and %s", gifts[index-1])
		} else {
			list += fmt.Sprintf("%s, ", gifts[index-1])
		}
	}
	return fmt.Sprintf("%s %s.", prefix, list)
}

func Song() string {
	var song []string
	for i := 1; i <= 12; i++ {
		song = append(song, Verse(i))
	}
	return strings.Join(song, "\n")
}
