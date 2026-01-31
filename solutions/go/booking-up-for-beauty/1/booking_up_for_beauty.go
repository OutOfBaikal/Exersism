package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"  // формат для строки "July 25, 2019 13:45:00"
    t, _ := time.Parse(layout, date)
    curTime := time.Now()
    return t.Before(curTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"  // формат для строки "Thursday, July 25, 2019 13:45:00"
    t, _ := time.Parse(layout, date)
    hour := t.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout, date)
    weekday := t.Format("Monday")
    month := t.Format("January")
    day := t.Format("2")
    year := t.Format("2006")
    hour := t.Format("15")
    minute := t.Format("04")
    desc := fmt.Sprintf("You have an appointment on %s, %s %s, %s, at %s:%s.", weekday, month, day, year, hour, minute)
    return desc
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	openingDate := time.Date(2019, 9, 15, 0, 0, 0, 0, time.UTC)
	
	// Get the current year
	currentYear := time.Now().Year()
	
	// Create a new date with the current year and the same month and day as the opening date
	anniversaryDate := time.Date(currentYear, openingDate.Month(), openingDate.Day(), 0, 0, 0, 0, time.UTC)
	
	return anniversaryDate
}
