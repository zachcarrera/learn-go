package resistorcolortrio

import (
	"math"
	"strconv"
	"strings"
)

var (
	resistors = map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	units = map[int]string{
		0: "ohms",
		1: "kiloohms",
		2: "megaohms",
		3: "gigaohms",
	}
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	if len(colors) < 3 {
		return ""
	}
	var sb strings.Builder
	base := resistors[colors[0]]*10 + resistors[colors[1]]
	numberOfZeros := resistors[colors[2]]
	multiplier := 1 * int(math.Pow10(numberOfZeros))
	base *= multiplier
	count := 0

	for len(strconv.Itoa(base)) > 3 {
		base /= 1000
		count++
	}
	sb.WriteString(strconv.Itoa(base))
	sb.WriteRune(' ')
	sb.WriteString(units[count])
	return sb.String()
}
