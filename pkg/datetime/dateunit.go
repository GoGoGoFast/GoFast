package datetime

// DateUnit represents a unit of time
type DateUnit int

const (
	Millisecond DateUnit = 1
	Second      DateUnit = 1000 * Millisecond
	Minute      DateUnit = 60 * Second
	Hour        DateUnit = 60 * Minute
	Day         DateUnit = 24 * Hour
	Week        DateUnit = 7 * Day
)

// GetMillis returns the number of milliseconds in the DateUnit
func (du DateUnit) GetMillis() int64 {
	return int64(du)
}
