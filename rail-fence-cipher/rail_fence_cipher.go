package railfence

import "strings"

func Encode(message string, rails int) string {
	newRails := make([]string, rails)
	i, sign := 0, 1
	for _, char := range message {
		newRails[i] += string(char)
		if cycle := i + sign; cycle == -1 || cycle == rails {
			sign *= -1
		}
		i += sign
	}
	return strings.Join(newRails, "")
}

func Decode(message string, rails int) string {
	messageRunes := []rune(message)
	decoded := make([]rune, len(messageRunes))

	var rails2 [][]int
	rail, dRail := 0, 1
	for i := 0; i < len(messageRunes); i++ {
		if i < rails {
			rails2 = append(rails2, []int{})
		}
		rails2[rail] = append(rails2[rail], i)

		if cycle := rail + dRail; cycle == -1 || cycle == rails {
			dRail *= -1
		}
		rail += dRail
	}
	var positions []int
	for _, slice := range rails2 {
		positions = append(positions, slice...)
	}

	for i, v := range positions {
		decoded[v] = messageRunes[i]

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
