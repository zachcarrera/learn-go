package queenattack

import (
	"errors"
	"regexp"
)

var errInvalidPosition = errors.New("Invalid queen position")

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {

	exp := regexp.MustCompile(`^[a-h][1-8]$`)

	if !exp.MatchString(whitePosition) || !exp.MatchString(blackPosition) || whitePosition == blackPosition {
		return false, errInvalidPosition
	}

	// wheck if queens are in the same row or column
	if whitePosition[0] == blackPosition[0] || whitePosition[1] == blackPosition[1] {
		return true, nil
	}

	// check if queens are aligned diagonally
	if whitePosition[0]-blackPosition[0] == whitePosition[1]-blackPosition[1] || whitePosition[0]-blackPosition[0] == blackPosition[1]-whitePosition[1] {
		return true, nil
	}

	return false, nil
}
