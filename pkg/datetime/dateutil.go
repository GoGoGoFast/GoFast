// Package datetime provides utility functions for handling date and time operations.
// This package simplifies parsing, formatting, and computing date and time values.
//
// Package datetime 提供了处理日期和时间操作的实用函数。
// 该包简化了解析、格式化和计算日期和时间值的过程。
package datetime

import (
	"errors"
	"fmt"
	"time"
)

// ParseDateString parses a date string into a time.Time object.
// The layout parameter should be a Go time layout string.
//
// ParseDateString 将日期字符串解析为 time.Time 对象。
// layout 参数应为 Go 时间布局字符串。
//
// Parameters:
// - dateStr: the date string to parse (要解析的日期字符串)
// - layout: the layout string to use for parsing (用于解析的布局字符串)
//
// Returns:
// - time.Time: the parsed time.Time object (解析后的 time.Time 对象)
// - error: if an error occurs during parsing (如果解析时发生错误)
func ParseDateString(dateStr, layout string) (time.Time, error) {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

// FormatDate formats a time.Time object into a string based on the provided layout.
//
// FormatDate 根据提供的布局格式化 time.Time 对象为字符串。
//
// Parameters:
// - date: the time.Time object to format (要格式化的 time.Time 对象)
// - layout: the layout string to use for formatting (用于格式化的布局字符串)
//
// Returns:
// - string: the formatted date string (格式化后的日期字符串)
func FormatDate(date time.Time, layout string) string {
	return date.Format(layout)
}

// GetDatePart extracts a specific part of the date (year, month, day, hour, minute, second).
//
// GetDatePart 提取日期的特定部分（年、月、日、时、分、秒）。
//
// Parameters:
// - date: the time.Time object to extract from (要提取的 time.Time 对象)
// - part: the part of the date to extract ("year", "month", "day", "hour", "minute", "second") (要提取的日期部分)
//
// Returns:
// - int: the extracted part of the date (提取的日期部分)
// - error: if an invalid part is specified (如果指定了无效的部分)
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
// GetStartOfDay 返回给定日期的开始时间。
//
// Parameters:
// - date: the time.Time object to get the start of the day for (要获取开始时间的 time.Time 对象)
//
// Returns:
// - time.Time: the start time of the day (一天的开始时间)
func GetStartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// GetEndOfDay returns the end time of the day for the given date.
//
// GetEndOfDay 返回给定日期的结束时间。
//
// Parameters:
// - date: the time.Time object to get the end of the day for (要获取结束时间的 time.Time 对象)
//
// Returns:
// - time.Time: the end time of the day (一天的结束时间)
func GetEndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())
}

// DateAdd adds a duration to the given date and returns the new date.
//
// DateAdd 向给定日期添加一个持续时间并返回新日期。
//
// Parameters:
// - date: the time.Time object to add the duration to (要添加持续时间的 time.Time 对象)
// - duration: the duration to add (要添加的持续时间)
//
// Returns:
// - time.Time: the new date with the added duration (添加持续时间后的新日期)
func DateAdd(date time.Time, duration time.Duration) time.Time {
	return date.Add(duration)
}

// DateDiff calculates the difference between two dates and returns the duration.
//
// DateDiff 计算两个日期之间的差异并返回持续时间。
//
// Parameters:
// - date1: the first date (第一个日期)
// - date2: the second date (第二个日期)
//
// Returns:
// - time.Duration: the duration between the two dates (两个日期之间的持续时间)
func DateDiff(date1, date2 time.Time) time.Duration {
	return date2.Sub(date1)
}

// FormatDuration formats a duration into a human-readable string.
//
// FormatDuration 将持续时间格式化为人类可读的字符串。
//
// Parameters:
// - duration: the time.Duration object to format (要格式化的 time.Duration 对象)
//
// Returns:
// - string: the formatted duration string (格式化后的持续时间字符串)
func FormatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// ZodiacSign returns the zodiac sign for the given date.
//
// ZodiacSign 返回给定日期的星座。
//
// Parameters:
// - date: the time.Time object to get the zodiac sign for (要获取星座的 time.Time 对象)
//
// Returns:
// - string: the zodiac sign (星座)
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
// ChineseZodiac 返回给定年份的中国生肖。
//
// Parameters:
// - year: the year to get the Chinese zodiac sign for (要获取中国生肖的年份)
//
// Returns:
// - string: the Chinese zodiac sign (中国生肖)
func ChineseZodiac(year int) string {
	zodiacs := []string{
		"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse", "Goat", "Monkey", "Rooster", "Dog", "Pig",
	}
	return zodiacs[year%12]
}

// DateRange generates a slice of dates between the start and end dates, inclusive.
//
// DateRange 生成开始日期和结束日期之间的日期切片，包括起始和结束日期。
//
// Parameters:
// - start: the start date (开始日期)
// - end: the end date (结束日期)
//
// Returns:
// - []time.Time: a slice of dates between the start and end dates (开始日期和结束日期之间的日期切片)
// - error: if the end date is before the start date (如果结束日期早于开始日期)
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
// IsLeapYear 检查给定年份是否为闰年。
//
// Parameters:
// - year: the year to check (要检查的年份)
//
// Returns:
// - bool: true if the year is a leap year, false otherwise (如果是闰年则返回 true，否则返回 false)
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
// DaysInMonth 返回特定年份中某个月的天数。
//
// Parameters:
// - year: the year (年份)
// - month: the month (月份)
//
// Returns:
// - int: the number of days in the month (该月的天数)
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
// AddMonths 向日期添加给定的月份数并返回新日期。
//
// Parameters:
// - date: the time.Time object to add months to (要添加月份的 time.Time 对象)
// - months: the number of months to add (要添加的月份数)
//
// Returns:
// - time.Time: the new date with the added months (添加月份后的新日期)
func AddMonths(date time.Time, months int) time.Time {
	return date.AddDate(0, months, 0)
}

// AddYears adds a given number of years to a date and returns the new date.
//
// AddYears 向日期添加给定的年份数并返回新日期。
//
// Parameters:
// - date: the time.Time object to add years to (要添加年份的 time.Time 对象)
// - years: the number of years to add (要添加的年份数)
//
// Returns:
// - time.Time: the new date with the added years (添加年份后的新日期)
func AddYears(date time.Time, years int) time.Time {
	return date.AddDate(years, 0, 0)
}

// GetCurrentTime returns the current local time.
//
// GetCurrentTime 返回当前本地时间。
//
// Returns:
// - time.Time: the current local time (当前本地时间)
func GetCurrentTime() time.Time {
	return time.Now()
}

// GetCurrentUTCTime returns the current UTC time.
//
// GetCurrentUTCTime 返回当前的 UTC 时间。
//
// Returns:
// - time.Time: the current UTC time (当前的 UTC 时间)
func GetCurrentUTCTime() time.Time {
	return time.Now().UTC()
}

// IsWeekend checks if a given date falls on a weekend.
//
// IsWeekend 检查给定日期是否在周末。
//
// Parameters:
// - date: the time.Time object to check (要检查的 time.Time 对象)
//
// Returns:
// - bool: true if the date is on a weekend, false otherwise (如果日期在周末则返回 true，否则返回 false)
func IsWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday checks if a given date falls on a weekday.
//
// IsWeekday 检查给定日期是否在工作日。
//
// Parameters:
// - date: the time.Time object to check (要检查的 time.Time 对象)
//
// Returns:
// - bool: true if the date is on a weekday, false otherwise (如果日期在工作日则返回 true，否则返回 false)
func IsWeekday(date time.Time) bool {
	weekday := date.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}

// DaysBetween calculates the number of days between two dates.
//
// DaysBetween 计算两个日期之间的天数。
//
// Parameters:
// - start: the start date (开始日期)
// - end: the end date (结束日期)
//
// Returns:
// - int: the number of days between the two dates (两个日期之间的天数)
func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// WeeksBetween calculates the number of weeks between two dates.
//
// WeeksBetween 计算两个日期之间的周数。
//
// Parameters:
// - start: the start date (开始日期)
// - end: the end date (结束日期)
//
// Returns:
// - int: the number of weeks between the two dates (两个日期之间的周数)
func WeeksBetween(start, end time.Time) int {
	return DaysBetween(start, end) / 7
}

// MonthsBetween calculates the number of months between two dates.
//
// MonthsBetween 计算两个日期之间的月数。
//
// Parameters:
// - start: the start date (开始日期)
// - end: the end date (结束日期)
//
// Returns:
// - int: the number of months between the two dates (两个日期之间的月数)
func MonthsBetween(start, end time.Time) int {
	years := end.Year() - start.Year()
	months := int(end.Month() - start.Month())
	return years*12 + months
}

// YearsBetween calculates the number of years between two dates.
//
// YearsBetween 计算两个日期之间的年数。
//
// Parameters:
// - start: the start date (开始日期)
// - end: the end date (结束日期)
//
// Returns:
// - int: the number of years between the two dates (两个日期之间的年数)
func YearsBetween(start, end time.Time) int {
	return end.Year() - start.Year()
}
