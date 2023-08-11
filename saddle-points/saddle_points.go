package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix [][]int

type Pair [2]int

func New(s string) (*Matrix, error) {
	if s == "" {
		return &Matrix{}, nil
	}
	var matrix Matrix
	for _, line := range strings.Split(s, "\n") {
		var row []int
		for _, col := range strings.Split(line, " ") {
			num, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	panic("Please implement the Saddle function")
}

func (m *Matrix) isMin(x, y int) bool {
	min := math.MaxInt
	for _, row := range *m {
		if row[y] < min {
			min = row[y]
		}
	}
	return (*m)[x][y] == min
}
