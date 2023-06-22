package romannumerals

import "errors"

type RomanConversion struct {
	arabic int
	roman  string
}

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("input must be between 1 and 3999")
	}

	dictionary := []RomanConversion{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var output string
	for _, conversion := range dictionary {
		for input >= conversion.arabic {
			output += conversion.roman
			input -= conversion.arabic
		}
	}
	return output, nil
}
