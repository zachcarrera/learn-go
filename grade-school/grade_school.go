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
	(*s)[student] = grade
}

func (s *School) Grade(level int) []string {
	var grade []string
	for studentName, gradeLevel := range *s {
		if gradeLevel == level {
			grade = append(grade, studentName)
		}
	}
	return grade
}

func (s *School) Enrollment() []Grade {
	panic("Please implement the Enrollment function")
}
