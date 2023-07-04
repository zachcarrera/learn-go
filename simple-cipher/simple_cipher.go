package cipher

import (
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
	encoded := make([]rune, utf8.RuneCountInString(input))
	for i, char := range input {
		encoded[i] = char + rune(c.distance)
	}
	return string(encoded)
}

func (c shift) Decode(input string) string {
	decoded := make([]rune, utf8.RuneCountInString(input))
	for i, char := range input {
		decoded[i] = char - rune(c.distance)
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
