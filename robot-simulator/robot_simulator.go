package robot

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

func StartRobot3(name, script string, action chan Action3, log chan string) {
	panic("Please implement the StartRobot3 function")
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	panic("Please implement the Room3 function")
}
