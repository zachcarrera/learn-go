package tournament

import (
	"fmt"
	"io"
)

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

func (t *teamRecord) String() string {
	return fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n", t.name, t.played, t.wins, t.draws, t.losses, t.points())
}

func Tally(reader io.Reader, writer io.Writer) error {
	panic("Please implement the Tally function")
}

func getOrCreateTeam(teams map[string]*teamRecord, name string) *teamRecord {
	team, ok := teams[name]
	if !ok {
		team = &teamRecord{name: name}
		teams[name] = team
	}
	return team
}
