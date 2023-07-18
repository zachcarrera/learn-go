package sublist

import "reflect"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	length1, length2 := len(l1), len(l2)

	switch {
	case length1 == length2 && reflect.DeepEqual(l1, l2):
		return RelationEqual
	case length1 > length2:
		opposite := Sublist(l2, l1)
		if opposite == RelationSublist {
			return RelationSuperlist
		}
	default:
		for i := 0; i <= length2-length1; i++ {
			if reflect.DeepEqual(l1, l2[i:i+length1]) {
				return RelationSublist
			}
		}
	}
	return RelationUnequal
}
