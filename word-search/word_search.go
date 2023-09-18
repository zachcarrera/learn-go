package wordsearch

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

func isOutOfBounds(x, y int, puzzle []string) bool {
	return x < 0 || y < 0 || x >= len(puzzle) || y >= len(puzzle[0])
}
