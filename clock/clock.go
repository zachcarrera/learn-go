package clock

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	panic("Please implement the Add function")
}

func (c Clock) Subtract(m int) Clock {
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	panic("Please implement the String function")
}
