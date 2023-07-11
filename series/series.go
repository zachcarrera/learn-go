package series

func All(n int, s string) []string {
	if n > len(s) {
		return []string{}
	}
	subStrings := make([]string, len(s)-n+1)
	characters := []rune(s)
	for i := 0; i < len(characters)-n+1; i++ {
		subString := make([]rune, n)
		for j := 0; j < n; j++ {
			subString[j] = characters[i+j]
		}
		subStrings[i] = string(subString)
	}
	return subStrings
}

func UnsafeFirst(n int, s string) string {
	panic("Please implement the UnsafeFirst function")
}
