package thefarm

import "errors"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	amount, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			if amount < 0 {
				return 0.0, errors.New("negative fodder")
			}
			return amount * 2 / float64(cows), nil
		}
		return 0, err
	}

	if amount < 0 {
		return 0.0, errors.New("negative fodder")
	}

	return amount / float64(cows), nil
}
