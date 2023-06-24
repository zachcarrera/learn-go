package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey %= 26
	rotated := ""
	for _, char := range plain {
		if !unicode.IsLetter(char) {
			rotated += string(char)
			continue
		}
		newChar := char + rune(shiftKey)
		if char <= 'Z' && newChar > 'Z' || (char <= 'z' && newChar > 'z') {
			newChar -= 26
		}
		rotated += string(newChar)
	}
	return rotated
}
