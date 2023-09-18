package wordsearch

import "fmt"

type direction struct {
	dx int
	dy int
}

var directions = []direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	panic("Please implement the Solve function")
}

func solveWord(word string, puzzle []string) ([2][2]int, error) {
	for startX, row := range puzzle {
		for startY := range row {
			for _, direction := range directions {
				completed := true
				x, y := startX, startY
				for _, char := range word {
					if isOutOfBounds(x, y, puzzle) || puzzle[x][y] != byte(char) {
						completed = false
						break
					}
					x += direction.dx
					y += direction.dy
				}
				if completed {
					return [2][2]int{{startY, startX}, {y - direction.dy, x - direction.dx}}, nil
				}
			}
		}
	}
	return [2][2]int{}, fmt.Errorf("word %s was not found in the puzzle", word)
}

func isOutOfBounds(x, y int, puzzle []string) bool {
	return x < 0 || y < 0 || x >= len(puzzle) || y >= len(puzzle[0])
}
