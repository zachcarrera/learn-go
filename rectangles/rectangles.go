package rectangles

type Corner [2]int

func Count(diagram []string) int {
	panic("Please implement the Count function")
}

func findCorners(diagram []string) []Corner {
	var corners []Corner
	for y, row := range diagram {
		for x, char := range row {
			if char == '+' {
				corners = append(corners, [2]int{x, y})
			}
		}
	}
	return corners
}
