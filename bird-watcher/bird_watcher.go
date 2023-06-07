package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, count := range birdsPerDay {
		sum += count
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	sum := 0

	lowerBound := 7 * (week - 1)
	upperBound := lowerBound + 7

	for i := lowerBound; i < upperBound; i++ {
		sum += birdsPerDay[i]
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
<<<<<<< HEAD

=======
>>>>>>> 40997b3d1dbea1fd776732fb0c4593e63c569dca
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
