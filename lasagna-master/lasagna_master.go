package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}
	return prepTimePerLayer * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	return 0, 0.0
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	return
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	return quantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
