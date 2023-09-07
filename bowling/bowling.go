package bowling

// Define the Game type here.
type Game struct {
	rolls         []int
	isSecondThrow bool
}

func NewGame() *Game {
	return &Game{rolls: make([]int, 0, 21)}
}

func (g *Game) Roll(pins int) error {
	panic("Please implement the Roll function")
}

func (g *Game) Score() (int, error) {
	panic("Please implement the Score function")
}
