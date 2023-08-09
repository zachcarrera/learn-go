package scale

import "strings"

func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmm"
	}
	interval += "m"
	notes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "e", "f", "bb", "eb":
		notes = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}
	var index int
	for i, note := range notes {
		if strings.EqualFold(note, tonic) {
			index = i
			break
		}
	}
	scale := make([]string, len(interval))
	for i, step := range interval {
		scale[i] = notes[index%len(notes)]
		switch step {
		case 'm':
			index += 1
		case 'M':
			index += 2
		case 'A':
			index += 3
		}
	}
	return scale
}
