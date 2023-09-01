package spiralmatrix

func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	row, col := 0, 0
	rowDirection, colDirection := 0, 1
	for i := 1; i <= size*size; i++ {
		matrix[row][col] = i

		// check when to switch directions, at the bounds of the matrix or when at a location that was already visited
		if row+rowDirection < 0 || size <= row+rowDirection ||
			col+colDirection < 0 || size <= col+colDirection ||
			matrix[row+rowDirection][col+colDirection] != 0 {
			rowDirection, colDirection = colDirection, -rowDirection
		}
		row, col = row+rowDirection, col+colDirection
	}
	return matrix
}
