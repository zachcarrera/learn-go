package bowling

// Define the Game type here.
type Game struct {
	rolls         []int
	isSecondThrow bool
}

func NewGame() *Game {
	panic("Please implement the NewGame function")
}

func (g *Game) Roll(pins int) error {
	panic("Please implement the Roll function")
}

func (g *Game) Score() (int, error) {
	panic("Please implement the Score function")
}
