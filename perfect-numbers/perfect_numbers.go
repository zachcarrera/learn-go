package perfect

import "errors"

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient Classification = 1 + iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("Input can only be positive")

func Classify(n int64) (Classification, error) {
	panic("Please implement the Classify function")
}
