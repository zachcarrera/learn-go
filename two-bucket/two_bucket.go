package twobucket

import (
	"errors"
)

const (
	bucketOne = "one"
	bucketTwo = "two"
)

type bucket struct {
	name         string
	capacity     int
	currentLevel int
}

func (b *bucket) isEmpty() bool {
	return b.currentLevel == 0
}

func (b *bucket) isFull() bool {
	return b.currentLevel == b.capacity
}

func (b *bucket) empty() {
	b.currentLevel = 0
}

func (b *bucket) fill() {
	b.currentLevel = b.capacity
}

func (b *bucket) availableSpace() int {
	return b.capacity - b.currentLevel
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if startBucket != bucketOne && startBucket != bucketTwo {
		return "", 0, 0, errors.New("invalid start bucket")
	}
	if sizeBucketOne == 0 || sizeBucketTwo == 0 {
		return "", 0, 0, errors.New("invalid bucket size")
	}
	if goalAmount == 0 || goalAmount > sizeBucketOne && goalAmount > sizeBucketTwo {
		return "", 0, 0, errors.New("invalid goal amount")
	}

	bucketsGCD := gcd(sizeBucketOne, sizeBucketTwo)
	if bucketsGCD > 1 && goalAmount%bucketsGCD != 0 {
		return "", 0, 0, errors.New("no solution with buckets")
	}

	var movesMade int
	var bucketA, bucketB *bucket

	if startBucket == bucketOne {
		bucketA = &bucket{bucketOne, sizeBucketOne, 0}
		bucketB = &bucket{bucketTwo, sizeBucketTwo, 0}

	} else {
		bucketA = &bucket{bucketTwo, sizeBucketTwo, 0}
		bucketB = &bucket{bucketOne, sizeBucketOne, 0}
	}

	for bucketA.currentLevel != goalAmount && bucketB.currentLevel != goalAmount {
		switch {
		case bucketA.isEmpty():
			bucketA.fill()
		case bucketB.isFull():
			bucketB.empty()
		case bucketB.isEmpty() && bucketB.capacity == goalAmount:
			bucketB.fill()
		default:
			// fill bucket b with bucket a
			if bucketB.availableSpace() >= bucketA.currentLevel {
				bucketB.currentLevel += bucketA.currentLevel
				bucketA.empty()
			} else {
				bucketA.currentLevel -= bucketB.availableSpace()
				bucketB.fill()
			}

		}
		movesMade++
	}

	if bucketA.currentLevel == goalAmount {
		return bucketA.name, movesMade, bucketB.currentLevel, nil
	}

	return bucketB.name, movesMade, bucketA.currentLevel, nil
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
