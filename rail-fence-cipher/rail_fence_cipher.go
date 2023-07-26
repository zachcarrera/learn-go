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
	panic("Please implement the Decode function")
}
