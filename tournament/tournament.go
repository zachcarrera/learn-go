package tournament

import "io"

type teamRecord struct {
	name   string
	played int
	wins   int
	losses int
	draws  int
}

func (t *teamRecord) won() {
	t.played++
	t.wins++
}

func (t *teamRecord) lost() {
	t.played++
	t.losses++
}

func (t *teamRecord) draw() {
	t.played++
	t.draws++
}

func (t *teamRecord) points() int {
	return t.wins*3 + t.draws
}

func Tally(reader io.Reader, writer io.Writer) error {
	panic("Please implement the Tally function")
}
