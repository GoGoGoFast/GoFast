package datetime

import "time"

// DayOfWeek represents days of the week
type DayOfWeek int

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// DayOfWeekFromTime converts time.Weekday to DayOfWeek
func DayOfWeekFromTime(weekday time.Weekday) DayOfWeek {
	return DayOfWeek(weekday)
}
