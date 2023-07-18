package sublist

import "reflect"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) {
		if reflect.DeepEqual(l1, l2) {
			return RelationEqual
		} else {
			return RelationUnequal
		}
	return RelationUnequal
	panic("Please implement the Sublist function")
}
