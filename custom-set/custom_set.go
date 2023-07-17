package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	set := make(Set)
	for _, v := range l {
		set.Add(v)
	}
	return set
}

func (s Set) String() string {
	keys := make([]string, 0, len(s))
	for key := range s {
		key = `"` + key + `"`
		keys = append(keys, key)
	}
	return "{" + strings.Join(keys, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for key := range s1 {
		if s2.Has(key) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	intersection := make([]string, 0)
	for key := range s1 {
		if s2.Has(key) {
			intersection = append(intersection, key)
		}
	}
	for key := range s2 {
		if s1.Has(key) {
			intersection = append(intersection, key)
		}
	}
	return NewFromSlice(intersection)
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
