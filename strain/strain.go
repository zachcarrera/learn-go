package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var filtered []int
	for _, number := range i {
		if filter(number) {
			filtered = append(filtered, number)
		}
	}
	return filtered
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var filtered []int
	for _, number := range i {
		if !filter(number) {
			filtered = append(filtered, number)
		}
	}
	return filtered
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	panic("Please implement the Keep function")
}

func (s Strings) Keep(filter func(string) bool) Strings {
	panic("Please implement the Keep function")
}
