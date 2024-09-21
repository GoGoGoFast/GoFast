package datetime

// Season represents the four seasons
type Season int

const (
	Spring Season = iota + 1 // 1~3 months
	Summer                   // 4~6 months
	Autumn                   // 7~9 months
	Winter                   // 10~12 months
)

// GetMonths returns the months that belong to the season
func (s Season) GetMonths() []Month {
	switch s {
	case Spring:
		return []Month{January, February, March}
	case Summer:
		return []Month{April, May, June}
	case Autumn:
		return []Month{July, August, September}
	case Winter:
		return []Month{October, November, December}
	default:
		return []Month{}
	}
}
