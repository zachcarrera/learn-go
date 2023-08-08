package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) []string {
	var knownAllergies []string
	for i, allergen := range allergens {
		if allergies&(1<<i) != 0 {
			knownAllergies = append(knownAllergies, allergen)
		}
	}
	return knownAllergies
}

func AllergicTo(allergies uint, allergen string) bool {
	panic("Please implement the AllergicTo function")
}
