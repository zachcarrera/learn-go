package queenattack

import "errors"

var errInvalidPosition = errors.New("Invalid queen position")

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {

	if whitePosition == "" || blackPosition == "" {
		return false, errInvalidPosition

	}
	if whitePosition == blackPosition {
		return false, errInvalidPosition
	}

	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errInvalidPosition
	}

	if whitePosition[0] < 'a' || whitePosition[0] > 'h' || blackPosition[0] < 'a' || blackPosition[0] > 'h' {
		return false, errInvalidPosition
	}
	if whitePosition[1] < '1' || whitePosition[1] > '8' || blackPosition[1] < '1' || blackPosition[1] > '8' {
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
