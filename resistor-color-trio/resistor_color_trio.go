package resistorcolortrio

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
	panic("Implement the Label function")
}
