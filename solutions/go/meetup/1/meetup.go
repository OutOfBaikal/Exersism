package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int 

const (
    First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// Create a date for the first day of the specified month and year
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// Initialize a counter for occurrences of the specified weekday
	occurrence := 0

	// Iterate through the days of the month
	for d := firstOfMonth; d.Month() == month; d = d.AddDate(0, 0, 1) {
		if d.Weekday() == wDay {
			occurrence++
			switch wSched {
			case First:
				if occurrence == 1 {
					return d.Day()
				}
			case Second:
				if occurrence == 2 {
					return d.Day()
				}
			case Third:
				if occurrence == 3 {
					return d.Day()
				}
			case Fourth:
				if occurrence == 4 {
					return d.Day()
				}
			case Teenth:
				if d.Day() >= 13 && d.Day() <= 19 {
					return d.Day()
				}
			case Last:
				lastDay := d.Day()
				// Continue to the next day to find the last occurrence
				for d = d.AddDate(0, 0, 1); d.Month() == month; d = d.AddDate(0, 0, 1) {
					if d.Weekday() == wDay {
						lastDay = d.Day()
					}
				}
				return lastDay
			}
		}
	}

	// If no day is found (which shouldn't happen), return 0
	return 0
}
