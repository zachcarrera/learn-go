package matrix

import (
	"math"
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
	var pairs []Pair
	for x, row := range *m {
		for y := range row {
			if m.isMin(x, y) && m.isMax(x, y) {
				pairs = append(pairs, Pair{x + 1, y + 1})
			}
		}
	}
	return pairs
}

func (m *Matrix) isMax(x, y int) bool {
	max := math.MinInt
	for _, num := range (*m)[x] {
		if max < num {
			max = num
		}
	}
	return (*m)[x][y] == max
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
