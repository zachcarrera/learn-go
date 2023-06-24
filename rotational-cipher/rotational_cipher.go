package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey = shiftKey % 26
	rotated := ""
	for _, char := range plain {
		if unicode.IsUpper(char) {
			new := char + rune(shiftKey)
			if new > 90 {
				new -= 26
			}
			rotated += string(new)
		} else if unicode.IsLower(char) {
			new := char + rune(shiftKey)
			if new > 122 {
				new -= 26
			}
			rotated += string(new)
		} else {
			rotated += string(char)
		}
	}
	return rotated
}
