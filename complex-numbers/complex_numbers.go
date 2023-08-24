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
	panic("Please implement the Add method")
}

func (n1 Number) Subtract(n2 Number) Number {
	panic("Please implement the Subtract method")
}

func (n1 Number) Multiply(n2 Number) Number {
	panic("Please implement the Multiply method")
}

func (n Number) Times(factor float64) Number {
	panic("Please implement the Times method")
}

func (n1 Number) Divide(n2 Number) Number {
	panic("Please implement the Divide method")
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
