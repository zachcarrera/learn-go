package anagram

func Detect(subject string, candidates []string) []string {
	frequencyTable := make(map[rune]int)
	for _, char := range subject {
		frequencyTable[char]++
	}
	return nil
}
