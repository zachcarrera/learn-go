package school

import "sort"

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
	var enrollment []Grade
	for studentName, gradeLevel := range *s {
		gradeExists := false
		for i := range enrollment {
			if enrollment[i].level == gradeLevel {
				gradeExists = true
				enrollment[i].students = append(enrollment[i].students, studentName)
				break
			}
		}
		if !gradeExists {
			enrollment = append(enrollment, Grade{level: gradeLevel, students: []string{studentName}})
		}
	}
	sort.Slice(enrollment, func(i, j int) bool {
		return enrollment[i].level < enrollment[j].level
	})
	for _, g := range enrollment {
		sort.Strings(g.students)
	}
	return enrollment
}
