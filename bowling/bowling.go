package bowling

import "errors"

// Define the Game type here.
type Game struct {
	rolls         []int
	isSecondThrow bool
}

func NewGame() *Game {
	return &Game{rolls: make([]int, 0, 21)}
}

func (g *Game) Roll(pins int) error {
	pinsKnockedDown := 0

	if g.isSecondThrow {
		pinsKnockedDown = g.rolls[len(g.rolls)-1]
	}

	if pins < 0 || pins+pinsKnockedDown > 10 {
		return errors.New("Pin count invalid")
	}

	// get the score and if the game is over return an error
	if _, err := g.Score(); err == nil {
		return errors.New("Cannot roll after a completed game")
	}

	g.rolls = append(g.rolls, pins)
	g.isSecondThrow = !g.isSecondThrow

	// strike was thrown
	if pins == 10 {
		g.isSecondThrow = false
	}
	return nil
}

func (g *Game) Score() (int, error) {
	var sum int
	for throwNumber, frame := 0, 0; frame < 10; throwNumber, frame = throwNumber+1, frame+1 {

		// check that both throws can be accessed from g.rolls
		if len(g.rolls) <= throwNumber+1 {
			return 0, errors.New("Game not complete")
		}

		one, two := g.rolls[throwNumber], g.rolls[throwNumber+1]
		sum += one + two

		// on a spare or strike add the bonus scores
		if one == 10 || one+two == 10 {
			if len(g.rolls) <= throwNumber+2 {
				return 0, errors.New("Game not complete")
			}
			sum += g.rolls[throwNumber+2]
		}

		// if the throw was not a strike have throwNumber increase by two by the next time the loop starts
		if one != 10 {
			throwNumber++
		}

	}
	return sum, nil
}
