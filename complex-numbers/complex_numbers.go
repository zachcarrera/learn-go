package complexnumbers

// Define the Number type here.

// Number represents a complex number where position 0 is a and position 1 is b
type Number [2]float64

func (n Number) Real() float64 {
	return n[0]
}

func (n Number) Imaginary() float64 {
	return n[1]
}

func (n1 Number) Add(n2 Number) Number {
	return [2]float64{n1[0] + n2[0], n1[1] + n2[1]}
}

func (n1 Number) Subtract(n2 Number) Number {
	return [2]float64{n1[0] - n2[0], n1[1] - n2[1]}
}

func (n1 Number) Multiply(n2 Number) Number {
	return [2]float64{
		(n1[0] * n2[0]) - (n1[1] * n2[1]),
		n1[1]*n2[0] + n1[0]*n2[1],
	}
}

func (n Number) Times(factor float64) Number {
	for i := range n {
		n[i] *= factor
	}
	return n
}

func (n1 Number) Divide(n2 Number) Number {
	return [2]float64{
		((n1[0] * n2[0]) + (n1[1] * n2[1])) / (math.Pow(n2[0], 2) + math.Pow(n2[1], 2)),
		(n1[1]*n2[0] - n1[0]*n2[1]) / (math.Pow(n2[0], 2) + math.Pow(n2[1], 2)),
	}
}

func (n Number) Conjugate() Number {
	panic("Please implement the Conjugate method")
}

func (n Number) Abs() float64 {
	panic("Please implement the Abs method")
}

func (n Number) Exp() Number {
	panic("Please implement the Exp method")
}
