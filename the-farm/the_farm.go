package thefarm

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			return amount * 1 / float64(cows), nil
		}
		return 0, err
	}

	return amount / float64(cows), nil
}
