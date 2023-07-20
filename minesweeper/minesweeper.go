package minesweeper

var mine = '*'

// Annotate returns an annotated board
func Annotate(board []string) []string {
	matrix := make([][]rune, len(board))
	for i, row := range board {
		matrix[i] = []rune(row)
	}
	for y, row := range matrix {
		for x := range row {
			if matrix[y][x] == mine {
				continue
			}
			var mineCount rune
			if y-1 >= 0 && x-1 >= 0 && matrix[y-1][x-1] == mine {
				mineCount++
			}
			if y-1 >= 0 && matrix[y-1][x] == mine {
				mineCount++
			}
			if y-1 >= 0 && x+1 < len(matrix[0]) && matrix[y-1][x+1] == mine {
				mineCount++
			}
			if x-1 >= 0 && matrix[y][x-1] == mine {
				mineCount++
			}
			if x+1 < len(matrix[0]) && matrix[y][x+1] == mine {
				mineCount++
			}
			if y+1 < len(matrix) && x-1 >= 0 && matrix[y+1][x-1] == mine {
				mineCount++
			}
			if y+1 < len(matrix) && matrix[y+1][x] == mine {
				mineCount++
			}
			if y+1 < len(matrix) && x+1 < len(matrix[0]) && matrix[y+1][x+1] == mine {
				mineCount++
			}

			if mineCount != 0 {
				matrix[y][x] = '0' + mineCount
			}
		}
	}
	boardStrings := make([]string, len(matrix))
	for i := range boardStrings {
		boardStrings[i] = string(matrix[i])
	}
	return boardStrings
}
