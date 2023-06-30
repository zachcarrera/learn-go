package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	h += m / 60
	m %= 60
	h %= 24

	if m < 0 {
		m += 60
		h -= 1
	}

	if h < 0 {
		h += 24
	}

	return Clock{
		hour:   h,
		minute: m,
	}
}

func (c Clock) Add(m int) Clock {
	c.minute += m
	c.hour += c.minute / 60
	c.minute %= 60
	c.hour %= 24
	return c
}

func (c Clock) Subtract(m int) Clock {
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
