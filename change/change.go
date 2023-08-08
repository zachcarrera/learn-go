package change

import "errors"

func Change(coins []int, target int) ([]int, error) {
	if target < 0 {
		return nil, errors.New("Cannot create negative change")
	}
	allPossibleChanges := make([][]int, target+1)
	allPossibleChanges[0] = []int{}
	for i := range allPossibleChanges {
		for _, coin := range coins {
			if i-coin >= 0 && allPossibleChanges[i-coin] != nil && (allPossibleChanges[i] == nil || len(allPossibleChanges[i-coin])+1 < len(allPossibleChanges[i])) {
				allPossibleChanges[i] = append([]int{coin}, allPossibleChanges[i-coin]...)
			}
		}
	}

	if allPossibleChanges[target] == nil {
		return nil, errors.New("Cannot create change")
	}

	return allPossibleChanges[target], nil
}
