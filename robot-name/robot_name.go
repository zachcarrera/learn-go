package robotname

// Define the Robot type here.
type Robot struct {
	name string
}

var names = map[string]bool{}

func (r *Robot) Name() (string, error) {
	panic("Please implement the Name function")
}

func (r *Robot) Reset() {
	r.name = ""
}
