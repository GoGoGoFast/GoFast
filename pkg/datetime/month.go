package datetime

import "time"

// Month represents months of the year
type Month int

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// GetLastDay returns the last day of the month
func (m Month) GetLastDay(leapYear bool) int {
	switch m {
	case January, March, May, July, August, October, December:
		return 31
	case April, June, September, November:
		return 30
	case February:
		if leapYear {
			return 29
		}
		return 28
	default:
		return 0
	}
}

// Of converts time.Month to Month
func Of(month time.Month) Month {
	return Month(month)
}
