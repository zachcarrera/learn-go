package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("input must be between 1 and 3999")
	}

	var output string
	switch input % 10 {
	case 1:
		output = "I"
	case 2:
		output = "II"
	case 3:
		output = "III"
	case 4:
		output = "IV"
	case 5:
		output = "V"
	case 6:
		output = "VI"
	case 7:
		output = "VII"
	case 8:
		output = "VIII"
	case 9:
		output = "IX"
	}

	switch (input / 10) % 10 {
	case 1:
		output = "X" + output
	case 2:
		output = "XX" + output
	case 3:
		output = "XXX" + output
	case 4:
		output = "XL" + output
	case 5:
		output = "L" + output
	case 6:
		output = "LX" + output
	case 7:
		output = "LXX" + output
	case 8:
		output = "LXXX" + output
	case 9:
		output = "XC" + output
	}

	switch (input / 100) % 10 {
	case 1:
		output = "C" + output
	case 2:
		output = "CC" + output
	case 3:
		output = "CCC" + output
	case 4:
		output = "CD" + output
	case 5:
		output = "D" + output
	case 6:
		output = "DC" + output
	case 7:
		output = "DCC" + output
	case 8:
		output = "DCCC" + output
	case 9:
		output = "CM" + output
	}

	switch (input / 1000) % 10 {
	case 1:
		output = "M" + output
	case 2:
		output = "MM" + output
	case 3:
		output = "MMM" + output
	}

	return output, nil
}
