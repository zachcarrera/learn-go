package twobucket

const (
	bucketOne = "one"
	bucketTwo = "two"
)

type bucket struct {
	name         string
	capacity     int
	currentLevel int
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	panic("Please implement the Solve function")
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
