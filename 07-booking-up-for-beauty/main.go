package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/02/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	rawTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(rawTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 1, 2006 15:04:05", date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	fmt.Println(Schedule("7/25/2019 13:45:00"))
	// => 2019-07-25 13:45:00 +0000 UTC

	fmt.Println(HasPassed("July 25, 2019 13:45:00"))
	// => true

	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	// => true

	fmt.Println(Description("7/25/2019 13:45:00"))
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."

	fmt.Println(AnniversaryDate())
	// => 2020-09-15 00:00:00 +0000 UTC
}
