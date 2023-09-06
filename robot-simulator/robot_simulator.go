package robot

import "fmt"

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const (
	N Dir = iota
	E
	S
	W
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
	case E:
		Step1Robot.X += 1
	case S:
		Step1Robot.Y -= 1
	case W:
		Step1Robot.X -= 1
	}
}

func (d Dir) String() string {
	return "NESW"[d : d+1]
}

// Step 2
// Define Action type here.
type Action byte

func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		action <- Action(c)
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case 'R':
			robot.Dir = (robot.Dir + 1) % 4
		case 'L':
			robot.Dir = (robot.Dir + 3) % 4
		case 'A':
			newPosition := robot.Pos
			switch robot.Dir {
			case N:
				newPosition.Northing += 1
			case E:
				newPosition.Easting += 1
			case S:
				newPosition.Northing -= 1
			case W:
				newPosition.Easting -= 1
			}
			if isInBounds(newPosition, extent) {
				robot.Pos = newPosition
			}
		}
	}
	report <- robot
}

func isInBounds(pos Pos, grid Rect) bool {
	return pos.Easting >= grid.Min.Easting && pos.Easting <= grid.Max.Easting &&
		pos.Northing >= grid.Min.Northing && pos.Northing <= grid.Max.Northing
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	name   string
	action Action
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	defer func() { action <- Action3{name: name, action: 0xFF} }()
	for _, command := range script {
		switch command {
		case 'R', 'L', 'A':
			action <- Action3{name: name, action: Action(command)}
		default:
			log <- fmt.Sprintf("invalid command %c in script", command)
			return
		}
	}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	defer func() { rep <- robots }()
	namePositions := make(map[string]int)
	filledPositions := make(map[Pos]int)
	for x, robot := range robots {

		// initialize robot
		if robot.Name == "" {
			log <- "Unnamed robot"
		}
		if _, ok := namePositions[robot.Name]; ok {
			log <- "Duplicate name"
		}
		namePositions[robot.Name] = x

		if !isInBounds(robot.Step2Robot.Pos, extent) {
			log <- "Robot placed outside the room"
			return
		}

		if _, ok := filledPositions[robot.Step2Robot.Pos]; ok {
			log <- "Position occupied"
			return
		}
		filledPositions[robot.Step2Robot.Pos] = x
	}
	completed := 0
	for a := range action {
		x, ok := namePositions[a.name]
		if !ok {
			log <- "Action by unknown robot"
			return
		}
		da := &robots[x].Step2Robot
		switch a.action {
		case 'R':
			da.Dir = (da.Dir + 1) % 4
		case 'L':
			da.Dir = (da.Dir + 3) % 4
		case 'A':
			newPosition := da.Pos
			switch da.Dir {
			case N:
				newPosition.Northing += 1
			case E:
				newPosition.Easting += 1
			case S:
				newPosition.Northing -= 1
			case W:
				newPosition.Easting -= 1
			}
			if !isInBounds(newPosition, extent) {
				log <- fmt.Sprintf("%s bumped into the wall", a.name)
				continue
			}
			if standingRobot, occupied := filledPositions[newPosition]; occupied {
				log <- fmt.Sprintf("%s bumped into %s", a.name, robots[standingRobot].Name)
				continue
			}
			delete(filledPositions, da.Pos)
			filledPositions[newPosition] = x
			da.Pos = newPosition
		case 0xFF:
			if completed++; completed == len(robots) {
				return
			}
		default:
			log <- "Undefined command"
			return
		}
	}
}
