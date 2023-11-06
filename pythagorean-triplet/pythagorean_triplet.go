package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	panic("Please implement the Range function")
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	panic("Please implement the Sum function")
}

func isPythagorean(a, b, c int) bool {
	return a*a+b*b == c*c
}
