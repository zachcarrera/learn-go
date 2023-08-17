package school

// Define the Grade and School types here.
type School map[string]int

type Grade struct {
	level    int
	students []string
}

func New() *School {
	school := make(School)
	return &school
}

func (s *School) Add(student string, grade int) {
	panic("Please implement the Add function")
}

func (s *School) Grade(level int) []string {
	panic("Please implement the Grade function")
}

func (s *School) Enrollment() []Grade {
	panic("Please implement the Enrollment function")
}
