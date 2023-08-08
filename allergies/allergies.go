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
	var position int
	for i, v := range allergens {
		if v == allergen {
			position = i
			break
		}
	}
	return allergies&(1<<position) != 0
}
