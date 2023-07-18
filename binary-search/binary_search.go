package binarysearch

func SearchInts(list []int, key int) int {
	low, high := 0, len(list)-1

	for high >= low {
		middle := (high + low) / 2
		if list[middle] == key {
			return middle
		} else if list[middle] < key {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return -1
}
