// Package datetime provides utility functions for handling date and time operations.
// This package simplifies parsing, formatting, and computing date and time values.
package datetime

import (
	"errors"
	"fmt"
	"time"
)

// ParseDateString parses a date string into a time.Time object.
// The layout parameter should be a Go time layout string.
//
// Parameters:
// - dateStr: the date string to parse
// - layout: the layout string to use for parsing
//
// Returns:
// - time.Time: the parsed time.Time object
// - error: if an error occurs during parsing
func ParseDateString(dateStr, layout string) (time.Time, error) {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

// FormatDate formats a time.Time object into a string based on the provided layout.
//
// Parameters:
// - date: the time.Time object to format
// - layout: the layout string to use for formatting
//
// Returns:
// - string: the formatted date string
func FormatDate(date time.Time, layout string) string {
	return date.Format(layout)
}

// GetDatePart extracts a specific part of the date (year, month, day, hour, minute, second).
//
// Parameters:
// - date: the time.Time object to extract from
// - part: the part of the date to extract ("year", "month", "day", "hour", "minute", "second")
//
// Returns:
// - int: the extracted part of the date
// - error: if an invalid part is specified
func GetDatePart(date time.Time, part string) (int, error) {
	switch part {
	case "year":
		return date.Year(), nil
	case "month":
		return int(date.Month()), nil
	case "day":
		return date.Day(), nil
	case "hour":
		return date.Hour(), nil
	case "minute":
		return date.Minute(), nil
	case "second":
		return date.Second(), nil
	default:
		return 0, errors.New("invalid date part")
	}
}

// GetStartOfDay returns the start time of the day for the given date.
//
// Parameters:
// - date: the time.Time object to get the start of the day for
//
// Returns:
// - time.Time: the start time of the day
func GetStartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// GetEndOfDay returns the end time of the day for the given date.
//
// Parameters:
// - date: the time.Time object to get the end of the day for
//
// Returns:
// - time.Time: the end time of the day
func GetEndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())
}

// DateAdd adds a duration to the given date and returns the new date.
//
// Parameters:
// - date: the time.Time object to add the duration to
// - duration: the duration to add
//
// Returns:
// - time.Time: the new date with the added duration
func DateAdd(date time.Time, duration time.Duration) time.Time {
	return date.Add(duration)
}

// DateDiff calculates the difference between two dates and returns the duration.
//
// Parameters:
// - date1: the first date
// - date2: the second date
//
// Returns:
// - time.Duration: the duration between the two dates
func DateDiff(date1, date2 time.Time) time.Duration {
	return date2.Sub(date1)
}

// FormatDuration formats a duration into a human-readable string.
//
// Parameters:
// - duration: the time.Duration object to format
//
// Returns:
// - string: the formatted duration string
func FormatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// ZodiacSign returns the zodiac sign for the given date.
//
// Parameters:
// - date: the time.Time object to get the zodiac sign for
//
// Returns:
// - string: the zodiac sign
func ZodiacSign(date time.Time) string {
	day := date.Day()
	month := date.Month()

	switch month {
	case time.January:
		if day < 20 {
			return "Capricorn"
		}
		return "Aquarius"
	case time.February:
		if day < 19 {
			return "Aquarius"
		}
		return "Pisces"
	case time.March:
		if day < 21 {
			return "Pisces"
		}
		return "Aries"
	case time.April:
		if day < 20 {
			return "Aries"
		}
		return "Taurus"
	case time.May:
		if day < 21 {
			return "Taurus"
		}
		return "Gemini"
	case time.June:
		if day < 21 {
			return "Gemini"
		}
		return "Cancer"
	case time.July:
		if day < 23 {
			return "Cancer"
		}
		return "Leo"
	case time.August:
		if day < 23 {
			return "Leo"
		}
		return "Virgo"
	case time.September:
		if day < 23 {
			return "Virgo"
		}
		return "Libra"
	case time.October:
		if day < 23 {
			return "Libra"
		}
		return "Scorpio"
	case time.November:
		if day < 22 {
			return "Scorpio"
		}
		return "Sagittarius"
	case time.December:
		if day < 22 {
			return "Sagittarius"
		}
		return "Capricorn"
	}
	return ""
}

// ChineseZodiac returns the Chinese zodiac sign for the given year.
//
// Parameters:
// - year: the year to get the Chinese zodiac sign for
//
// Returns:
// - string: the Chinese zodiac sign
func ChineseZodiac(year int) string {
	zodiacs := []string{
		"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse", "Goat", "Monkey", "Rooster", "Dog", "Pig",
	}
	return zodiacs[year%12]
}

// DateRange generates a slice of dates between the start and end dates, inclusive.
//
// Parameters:
// - start: the start date
// - end: the end date
//
// Returns:
// - []time.Time: a slice of dates between the start and end dates
// - error: if the end date is before the start date
func DateRange(start, end time.Time) ([]time.Time, error) {
	if end.Before(start) {
		return nil, errors.New("end date must be after start date")
	}

	var dates []time.Time
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}
	return dates, nil
}

// IsLeapYear checks if a given year is a leap year.
//
// Parameters:
// - year: the year to check
//
// Returns:
// - bool: true if the year is a leap year, false otherwise
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

// DaysInMonth returns the number of days in a given month of a specific year.
//
// Parameters:
// - year: the year
// - month: the month
//
// Returns:
// - int: the number of days in the month
func DaysInMonth(year int, month time.Month) int {
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

// AddMonths adds a given number of months to a date and returns the new date.
//
// Parameters:
// - date: the time.Time object to add months to
// - months: the number of months to add
//
// Returns:
// - time.Time: the new date with the added months
func AddMonths(date time.Time, months int) time.Time {
	return date.AddDate(0, months, 0)
}

// AddYears adds a given number of years to a date and returns the new date.
//
// Parameters:
// - date: the time.Time object to add years to
// - years: the number of years to add
//
// Returns:
// - time.Time: the new date with the added years
func AddYears(date time.Time, years int) time.Time {
	return date.AddDate(years, 0, 0)
}

// GetCurrentTime returns the current local time.
//
// Returns:
// - time.Time: the current local time
func GetCurrentTime() time.Time {
	return time.Now()
}

// GetCurrentUTCTime returns the current UTC time.
//
// Returns:
// - time.Time: the current UTC time
func GetCurrentUTCTime() time.Time {
	return time.Now().UTC()
}

// IsWeekend checks if a given date falls on a weekend.
//
// Parameters:
// - date: the time.Time object to check
//
// Returns:
// - bool: true if the date is on a weekend, false otherwise
func IsWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday checks if a given date falls on a weekday.
//
// Parameters:
// - date: the time.Time object to check
//
// Returns:
// - bool: true if the date is on a weekday, false otherwise
func IsWeekday(date time.Time) bool {
	weekday := date.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}

// DaysBetween calculates the number of days between two dates.
//
// Parameters:
// - start: the start date
// - end: the end date
//
// Returns:
// - int: the number of days between the two dates
func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// WeeksBetween calculates the number of weeks between two dates.
//
// Parameters:
// - start: the start date
// - end: the end date
//
// Returns:
// - int: the number of weeks between the two dates
func WeeksBetween(start, end time.Time) int {
	return DaysBetween(start, end) / 7
}

// MonthsBetween calculates the number of months between two dates.
//
// Parameters:
// - start: the start date
// - end: the end date
//
// Returns:
// - int: the number of months between the two dates
func MonthsBetween(start, end time.Time) int {
	years := end.Year() - start.Year()
	months := int(end.Month() - start.Month())
	return years*12 + months
}

// YearsBetween calculates the number of years between two dates.
//
// Parameters:
// - start: the start date
// - end: the end date
//
// Returns:
// - int: the number of years between the two dates
func YearsBetween(start, end time.Time) int {
	return end.Year() - start.Year()
}
