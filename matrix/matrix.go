package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {

	rows := strings.Split(s, "\n")
	m := make(Matrix, 0, len(rows))
	var rowLength int
	for i, row := range rows {
		num := strings.Split(strings.TrimSpace(row), " ")
		var numbers []int
		for _, v := range num {
			number, err := strconv.Atoi(v)
			if err != nil {
				return nil, errors.New("test error")
			}
			numbers = append(numbers, number)
		}
		if i == 0 {
			rowLength = len(numbers)
		} else {
			if len(numbers) != rowLength {
				return nil, errors.New("length error")
			}
		}
		m = append(m, numbers)
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	panic("Please implement the Cols function")
}

func (m Matrix) Rows() [][]int {
	panic("Please implement the Rows function")
}

func (m Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
