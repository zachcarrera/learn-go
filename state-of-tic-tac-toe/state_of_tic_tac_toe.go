package stateoftictactoe

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	panic("Please implement the StateOfTicTacToe function")
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
