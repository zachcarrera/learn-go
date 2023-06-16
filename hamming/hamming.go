package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences must match")
	}
	differences := 0
	bSlice := []rune(b)
	for i, char := range a {
		if char != bSlice[i] {
			differences += 1
		}
	}
	return differences, nil
}
