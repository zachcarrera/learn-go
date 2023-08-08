package phonenumber

import (
	"errors"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var cleanPhoneNumber strings.Builder
	for _, char := range phoneNumber {
		if unicode.IsDigit(char) {
			cleanPhoneNumber.WriteRune(char)
		}
	}
	cleanedDigits := cleanPhoneNumber.String()
	if len(cleanedDigits) == 11 && cleanedDigits[0] == '1' {
		cleanedDigits = cleanedDigits[1:]
	}

	if cleanedDigits[0] < '2' || cleanedDigits[3] < '2' || len(cleanedDigits) != 10 {
		return "", errors.New("invalid phone number")
	}

	return cleanedDigits, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	panic("Please implement the Format function")
}
