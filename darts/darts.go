package darts

import "math"

func Score(x, y float64) int {
	// find direct distance from center (0, 0)
	// use that distance to calculate points

	// a^2 + b^2 = c^2
	// c = sqrt(a^2 + b^2)
	distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	default:
		return 0
	}
}
