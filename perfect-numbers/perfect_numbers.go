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
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	aliquotSum := 1
	for i := 2; i < int(n); i++ {
		if int(n)%i == 0 {
			aliquotSum += i
		}
	}
	if n == 1 || aliquotSum < int(n) {
		return ClassificationDeficient, nil
	} else if aliquotSum == int(n) {
		return ClassificationPerfect, nil
	} else {
		return ClassificationAbundant, nil
	}
}
