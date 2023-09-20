package rectangles

type Corner [2]int

func Count(diagram []string) int {
	var rectangleCount int
	corners := findCorners(diagram)
	for _, c := range corners {
		for _, c2 := range corners {
			// check that the corners don't share an edge
			// then check that those corners make a valid rectangle
			if c[0] < c2[0] && c[1] < c2[1] && isConnectedRectangle(c, c2, diagram) {
				rectangleCount++
			}
		}
	}
	return rectangleCount
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

func isConnectedRectangle(c1, c2 Corner, diagram []string) bool {
	x1, y1 := c1[0], c1[1]
	x2, y2 := c2[0], c2[1]

	// check if the other two corners of a possible rectangle both contain '+'
	if diagram[y1][x2] != '+' || diagram[y2][x1] != '+' {
		return false
	}

	// check that the top and bottom edges contain valid chars, '-' or '+'
	for _, topEdgeChar := range diagram[y1][x1:x2] {
		for _, bottomEdgeChar := range diagram[y2][x1:x2] {
			if !isValidHorizontal(topEdgeChar) || !isValidHorizontal(bottomEdgeChar) {
				return false
			}
		}
	}

	// check that both vertical edges contain valid chars, '|' or '+'
	for _, row := range diagram[y1:y2] {
		leftEdgeChar := row[x1]
		rightEdgeChar := row[x2]
		if (leftEdgeChar != '|' && leftEdgeChar != '+') || (rightEdgeChar != '|' && rightEdgeChar != '+') {
			return false
		}
	}
	return true
}

func isValidHorizontal(char rune) bool {
	return char == '-' || char == '+'
}
