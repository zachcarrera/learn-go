package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Fifth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	switch wSched {
	case First:
		for wDay != date.Weekday() {
			date = date.AddDate(0, 0, 1)
		}
	case Second:
		date = date.AddDate(0, 0, 7)
		for wDay != date.Weekday() {
			date = date.AddDate(0, 0, 1)
		}
	case Third:
		date = date.AddDate(0, 0, 14)
		for wDay != date.Weekday() {
			date = date.AddDate(0, 0, 1)
		}
	case Fourth:
		date = date.AddDate(0, 0, 21)
		for wDay != date.Weekday() {
			date = date.AddDate(0, 0, 1)
		}
	case Last:
		date = time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local)
		date = date.AddDate(0, 0, -1)
		for wDay != date.Weekday() {
			date = date.AddDate(0, 0, -1)
		}
	case Teenth:
		date = time.Date(year, month, 13, 0, 0, 0, 0, time.Local)
		for wDay != date.Weekday() {
			date = date.AddDate(0, 0, 1)
		}
	}
	return date.Day()
}
