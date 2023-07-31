package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	freqChan := make(chan FreqMap)
	for _, text := range texts {
		go func(s string) {
			freqChan <- Frequency(s)
		}(text)
	}

	mainMap := make(FreqMap)
	for i := 0; i < len(texts); i++ {
		for k, v := range <-freqChan {
			mainMap[k] += v
		}
	}
	return mainMap
}
