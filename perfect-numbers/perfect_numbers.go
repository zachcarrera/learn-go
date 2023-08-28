package perfect

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient Classification = 1 + iota
	ClassificationPerfect
	ClassificationAbundant
)

func Classify(n int64) (Classification, error) {
	panic("Please implement the Classify function")
}
