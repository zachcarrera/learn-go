package twelve

import (
	"fmt"
	"strings"
)

var (
	gifts = []string{
		"twelve Drummers Drumming,",
		"eleven Pipers Piping,",
		"ten Lords-a-Leaping,",
		"nine Ladies Dancing,",
		"eight Maids-a-Milking,",
		"seven Swans-a-Swimming,",
		"six Geese-a-Laying,",
		"five Gold Rings,",
		"four Calling Birds,",
		"three French Hens,",
		"two Turtle Doves, and",
		"a Partridge in a Pear Tree.",
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
	list := strings.Join(gifts[12-i:], " ")
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", days[i-1], list)
}

func Song() string {
	song := make([]string, 12)
	for i := 1; i <= 12; i++ {
		song[i-1] = Verse(i)
	}
	return strings.Join(song, "\n")
}
