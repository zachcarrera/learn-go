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
	} else if len(l1) > len(l2) {
		if result := Sublist(l2, l1); result != "sublist" {
			return result
		}
		return RelationSuperlist
	} else {
		for i := 0; i <= len(l2)-len(l1); i++ {
			if reflect.DeepEqual(l1, l2[i:i+len(l1)]) {
				return RelationSublist
			}
		}
	}
	return RelationUnequal
}
