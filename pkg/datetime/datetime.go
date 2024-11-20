package datetime

import (
	"time"
)

// DateTime is a wrapper around time.Time to provide more utility methods
type DateTime struct {
	time.Time
}

// NewDateTime creates a new DateTime instance
//
// Parameters:
// - t: the time.Time object to wrap
//
// Returns:
// - *DateTime: the new DateTime instance
func NewDateTime(t time.Time) *DateTime {
	return &DateTime{Time: t}
}

// Now returns the current DateTime
//
// Returns:
// - *DateTime: the current DateTime
func Now() *DateTime {
	return NewDateTime(time.Now())
}

// Format formats the DateTime according to the given pattern
//
// Parameters:
// - pattern: the format pattern
//
// Returns:
// - string: the formatted DateTime string
func (dt *DateTime) Format(pattern string) string {
	return dt.Time.Format(pattern)
}

// AddDays adds the given number of days to the DateTime
//
// Parameters:
// - days: the number of days to add
//
// Returns:
// - *DateTime: the new DateTime with the added days
func (dt *DateTime) AddDays(days int) *DateTime {
	return NewDateTime(dt.Time.AddDate(0, 0, days))
}

// AddMonths adds the given number of months to the DateTime
//
// Parameters:
// - months: the number of months to add
//
// Returns:
// - *DateTime: the new DateTime with the added months
func (dt *DateTime) AddMonths(months int) *DateTime {
	return NewDateTime(dt.Time.AddDate(0, months, 0))
}

// AddYears adds the given number of years to the DateTime
//
// Parameters:
// - years: the number of years to add
//
// Returns:
// - *DateTime: the new DateTime with the added years
func (dt *DateTime) AddYears(years int) *DateTime {
	return NewDateTime(dt.Time.AddDate(years, 0, 0))
}

// Between calculates the duration between two DateTime instances
//
// Parameters:
// - other: the other DateTime instance to compare with
//
// Returns:
// - time.Duration: the duration between the two DateTime instances
func (dt *DateTime) Between(other *DateTime) time.Duration {
	return dt.Time.Sub(other.Time)
}
