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
	input = strings.ToLower(input)
	encoded := make([]rune, utf8.RuneCountInString(input))
	pos := 0
	for _, char := range input {
		if unicode.IsLower(char) {
			encoded[pos] = (char+rune(c.distance)-'a')%26 + 'a'
			if encoded[pos] < 'a' {
				encoded[pos] += 26
			}
			pos++
		}
	}
	return string(encoded[:pos])
}

func (c shift) Decode(input string) string {
	decoded := make([]rune, utf8.RuneCountInString(input))
	for i, char := range input {
		if unicode.IsLower(char) {
			decoded[i] = (char-rune(c.distance)-'a')%26 + 'a'
		}
		if decoded[i] < 'a' {
			decoded[i] += 26
		}
	}
	return string(decoded)
}

func NewVigenere(key string) Cipher {
	lower := true
	allA := true
	for _, char := range key {
		if !unicode.IsLower(char) {
			lower = false
		}
		if char != 'a' {
			allA = false
		}
	}

	if !lower || allA {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	encoded := make([]rune, utf8.RuneCountInString(input))
	pos := 0
	for _, char := range input {
		if unicode.IsLower(char) {
			shift := []rune(v.key)[pos%utf8.RuneCountInString(v.key)] - 'a'
			encoded[pos] = (char+shift-'a')%26 + 'a'
			pos++
		}
	}
	return string(encoded[:pos])
}

func (v vigenere) Decode(input string) string {
	decoded := make([]rune, utf8.RuneCountInString(input))
	for i, char := range input {
		shift := []rune(v.key)[i%utf8.RuneCountInString(v.key)] - 'a'
		decoded[i] = (char-shift-'a')%26 + 'a'
		if decoded[i] < 'a' {
			decoded[i] += 26
		}
	}
	return string(decoded)
}
