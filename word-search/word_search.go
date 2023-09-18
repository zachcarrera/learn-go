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
