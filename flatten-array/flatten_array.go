package flatten

import "reflect"

func Flatten(nested interface{}) []interface{} {
	flattened := make([]interface{}, 0)
	nestedTyped := nested.([]interface{})
	for _, v := range nestedTyped {
		if v == nil {
			continue
		} else if reflect.TypeOf(v).Kind() == reflect.Slice {
			flattened = append(flattened, Flatten(v)...)
		} else {
			flattened = append(flattened, v)
		}
	}
	return flattened
}
