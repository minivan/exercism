// Package clock lets us run clock arithmetic operations
package clock

import (
	"fmt"
)

// Clock represents the time using hours and minutes
type Clock struct {
	Hours   int
	Minutes int
}

// New returns a new Clock struct
func New(hour, minute int) Clock {
	return Clock{hour, minute}.Normalize()
}

// String returns a string respresentation of a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

// Normalize calculates the correct time representation from hours to minutes
func (c Clock) Normalize() Clock {
	if c.Minutes < 0 {
		// If the minutes are negative, we need to subtract an extra hour
		subtractedHours := (-c.Minutes / 60) + 1
		c.Minutes += subtractedHours * 60
		c.Hours -= subtractedHours
	}
	if c.Hours < 0 {
		c.Hours %= 24
		c.Hours += 24
	}
	if c.Minutes >= 60 {
		addedHours := c.Minutes / 60
		c.Minutes %= 60
		c.Hours += addedHours
	}
	if c.Hours >= 24 {
		c.Hours %= 24
	}
	return c
}

// Add adds minutes to a clock
func (c Clock) Add(minutes int) Clock {
	c.Minutes += minutes

	return c.Normalize()
}

// Subtract subtracts minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	c.Minutes -= minutes

	return c.Normalize()
}
