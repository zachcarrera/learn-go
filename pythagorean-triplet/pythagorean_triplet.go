package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var triplets []Triplet
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if isPythagorean(a, b, c) {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var triplets []Triplet
	max := p / 2
	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := p - a - b; isPythagorean(a, b, c) {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}

func isPythagorean(a, b, c int) bool {
	return a*a+b*b == c*c
}
