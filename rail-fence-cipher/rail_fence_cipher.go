package railfence

func Encode(message string, rails int) string {
	messageRunes := []rune(message)
	length := len(messageRunes)
	encoded := make([]rune, length)
	positions := makeRails(length, rails)

	for to, from := range positions {
		encoded[to] = messageRunes[from]
	}
	return string(encoded)
}

func Decode(message string, rails int) string {
	messageRunes := []rune(message)
	length := len(messageRunes)
	decoded := make([]rune, length)
	positions := makeRails(length, rails)

	for from, to := range positions {
		decoded[to] = messageRunes[from]
	}
	return string(decoded)
}

// makeRails returns an int slice that represents the index that a rune needs to be placed in a slice to encode or decode a cipher.
func makeRails(length, numRails int) []int {
	var rails [][]int
	rail, dRail := 0, 1
	for i := 0; i < length; i++ {
		if i < numRails {
			rails = append(rails, []int{})
		}
		rails[rail] = append(rails[rail], i)
		if cycle := rail + dRail; cycle == -1 || cycle == numRails {
			dRail *= -1
		}
		rail += dRail
	}
	var positions []int
	for _, slice := range rails {
		positions = append(positions, slice...)
	}
	return positions
}
