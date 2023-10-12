package stateoftictactoe

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	xCount, oCount := countPlayer(board, 'X'), countPlayer(board, 'O')
	if xCount != oCount && xCount != oCount+1 {
		return "", errors.New("out of turn play")
	}
	xWon, oWon := hasWon(board, 'X'), hasWon(board, 'O')
	switch {
	case xWon && oWon:
		return "", errors.New("game continued after a win")
	case xWon || oWon:
		return Win, nil
	case xCount+oCount == 9:
		return Draw, nil
	}
	return Ongoing, nil
}

func countPlayer(board []string, player rune) int {
	count := 0
	for _, row := range board {
		for _, square := range row {
			if square == player {
				count++
			}
		}
	}
	return count
}

func hasWon(board []string, player byte) bool {

	// check for a horizontal or vertical win
	for i := 0; i < len(board); i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) || (board[0][i] == player && board[1][i] == player && board[2][i] == player) {
			return true
		}
	}

	// check if there is a diagnal win
	return (board[0][0] == player && board[1][1] == player && board[2][2] == player) || (board[0][2] == player && board[1][1] == player && board[2][0] == player)
}
