package resistorcolor

// Colors should return the list of all colors.
func Colors() []string {
	return []string{
		"Black:",
		"Brown:",
		"Red:",
		"Orange:",
		"Yellow:",
		"Green:",
		"Blue:",
		"Violet:",
		"Grey:",
		"White:",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	panic("Please implement the ColorCode function")
}
