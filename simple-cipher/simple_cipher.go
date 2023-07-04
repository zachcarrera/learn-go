package cipher

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Define the shift and vigenere types here.
type shift struct {
	distance int
}
type vigenere struct {
	key string
}

// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	// TODO: fix wrap for negative shift
	input = strings.ToLower(input)
	encoded := make([]rune, utf8.RuneCountInString(input))
	pos := 0
	for _, char := range input {
		if unicode.IsLower(char) {
			encoded[pos] = (char+rune(c.distance)-'a')%26 + 'a'
			pos++
		}
	}
	return string(encoded[:pos])
}

func (c shift) Decode(input string) string {
	// TODO: Fix wrapping shifts for negative shifts
	decoded := make([]rune, utf8.RuneCountInString(input))
	for i, char := range input {
		if unicode.IsLower(char) {
			decoded[i] = (char-rune(c.distance)-'a')%26 + 'a'
		}
	}
	return string(decoded)
}

func NewVigenere(key string) Cipher {
	panic("Please implement the NewVigenere function")
}

func (v vigenere) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (v vigenere) Decode(input string) string {
	panic("Please implement the Decode function")
}