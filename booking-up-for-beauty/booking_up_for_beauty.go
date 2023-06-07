package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	return time.Since(parsedTime) > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)
	weekday := parsedTime.Weekday().String()
	month := parsedTime.Month().String()
	day := parsedTime.Day()
	year := parsedTime.Year()
	hour := parsedTime.Hour()
	minute := parsedTime.Minute()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", weekday, month, day, year, hour, minute)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
