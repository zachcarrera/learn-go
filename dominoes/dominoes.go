package dominoes

// Define the Domino type here.
type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 || len(input) == 1 && input[0][0] == input[0][1] {
		return input, true
	}
	currentDomino := input[0]
	for i, domino := range input {

		// skip the first domino
		if i == 0 {
			continue
		}

		// if a match can be made by flipping a domino then flip it
		if domino[1] == currentDomino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}
		if domino[0] == currentDomino[1] {
			// merge the matched pair into one domino
			// and try to match it with the rest of the dominoes
			newDomino := [2]int{currentDomino[0], domino[1]}
			newDominoes := append([]Domino{newDomino}, input[1:i]...)
			newDominoes = append(newDominoes, input[i+1:]...)
			oldSolution, ok := MakeChain(newDominoes)
			if ok {
				return append([]Domino{currentDomino, domino}, oldSolution[1:]...), true
			}
		}

	}
	return nil, false
}
