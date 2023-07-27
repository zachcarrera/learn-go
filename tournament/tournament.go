package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
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
	teams, err := readMatches(reader)
	if err != nil {
		return err
	}

	topRow := fmt.Sprintf("%-30s |%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	writer.Write([]byte(topRow))
	for _, team := range rankedTeams(teams) {
		writer.Write([]byte(team.String()))
	}
	return nil
}

func readMatches(reader io.Reader) (map[string]*teamRecord, error) {
	teams := make(map[string]*teamRecord)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" || text[0] == '#' {
			continue
		}

		tokens := strings.Split(text, ";")
		if len(tokens) != 3 {
			return nil, errors.New("Invalid match given")
		}

		switch tokens[2] {
		case "win":
			getOrCreateTeam(teams, tokens[0]).won()
			getOrCreateTeam(teams, tokens[1]).lost()
		case "loss":
			getOrCreateTeam(teams, tokens[0]).lost()
			getOrCreateTeam(teams, tokens[1]).won()
		case "draw":
			getOrCreateTeam(teams, tokens[0]).draw()
			getOrCreateTeam(teams, tokens[1]).draw()
		default:
			return nil, errors.New("Invalid match result")
		}
	}
	return teams, nil
}

func rankedTeams(teams map[string]*teamRecord) []*teamRecord {
	rankedTeams := make([]*teamRecord, 0, len(teams))
	for _, team := range teams {
		rankedTeams = append(rankedTeams, team)
	}

	sort.Slice(rankedTeams, func(a, b int) bool {
		if rankedTeams[a].points() > rankedTeams[b].points() {
			return true
		}
		if rankedTeams[a].points() < rankedTeams[b].points() {
			return false
		}

		return strings.Compare(rankedTeams[a].name, rankedTeams[b].name) < 0
	})
	return rankedTeams
}

func getOrCreateTeam(teams map[string]*teamRecord, name string) *teamRecord {
	team, ok := teams[name]
	if !ok {
		team = &teamRecord{name: name}
		teams[name] = team
	}
	return team
}
