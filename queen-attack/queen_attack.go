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

	whiteXY := []int{
		int(whitePosition[0] - 'a'),
		int(whitePosition[1] - '1'),
	}
	blackXY := []int{
		int(blackPosition[0] - 'a'),
		int(blackPosition[1] - '1'),
	}

	if whiteXY[0] == blackXY[0] {
		return true, nil
	}

	if whiteXY[1] == blackXY[1] {
		return true, nil
	}

	if whiteXY[0]-blackXY[0] == whiteXY[1]-blackXY[1] || whiteXY[0]-blackXY[0] == -1*(whiteXY[1]-blackXY[1]) {
		return true, nil
	}

	return false, nil
}
