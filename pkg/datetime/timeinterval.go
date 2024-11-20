// Package datetime provides utility functions for handling time intervals and durations.
// This package simplifies measuring the execution time of code.
package datetime

import (
	"time"
)

// TimeInterval is a simple timer class used to calculate the duration of code execution
type TimeInterval struct {
	start time.Time
}

// NewTimeInterval creates a new TimeInterval instance and starts the timer
//
// Returns:
// - *TimeInterval: a new TimeInterval instance
func NewTimeInterval() *TimeInterval {
	return &TimeInterval{start: time.Now()}
}

// Elapsed returns the elapsed time in the specified unit
//
// Parameters:
// - unit: the unit of time to return
//
// Returns:
// - int64: the elapsed time in the specified unit
func (ti *TimeInterval) Elapsed(unit DateUnit) int64 {
	duration := time.Since(ti.start)
	return duration.Nanoseconds() / (unit.GetMillis() * 1e6)
}

// ElapsedMillis returns the elapsed time in milliseconds
//
// Returns:
// - int64: the elapsed time in milliseconds
func (ti *TimeInterval) ElapsedMillis() int64 {
	return ti.Elapsed(Millisecond)
}

// ElapsedSeconds returns the elapsed time in seconds
//
// Returns:
// - int64: the elapsed time in seconds
func (ti *TimeInterval) ElapsedSeconds() int64 {
	return ti.Elapsed(Second)
}

// ElapsedMinutes returns the elapsed time in minutes
//
// Returns:
// - int64: the elapsed time in minutes
func (ti *TimeInterval) ElapsedMinutes() int64 {
	return ti.Elapsed(Minute)
}

// ElapsedHours returns the elapsed time in hours
//
// # ElapsedHours
//
// Returns:
// - int64: the elapsed time in hours
func (ti *TimeInterval) ElapsedHours() int64 {
	return ti.Elapsed(Hour)
}

// ElapsedDays returns the elapsed time in days
//
// # ElapsedDays
//
// Returns:
// - int64: the elapsed time in days
func (ti *TimeInterval) ElapsedDays() int64 {
	return ti.Elapsed(Day)
}

// ElapsedWeeks returns the elapsed time in weeks
//
// # ElapsedWeeks
//
// Returns:
// - int64: the elapsed time in weeks
func (ti *TimeInterval) ElapsedWeeks() int64 {
	return ti.Elapsed(Week)
}
