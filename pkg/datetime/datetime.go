package datetime

import (
	"time"
)

// DateTime is a wrapper around time.Time to provide more utility methods
// DateTime 是 time.Time 的包装器，提供更多实用方法
type DateTime struct {
	time.Time
}

// NewDateTime creates a new DateTime instance
// NewDateTime 创建一个新的 DateTime 实例
//
// Parameters:
// - t: the time.Time object to wrap (要包装的 time.Time 对象)
//
// Returns:
// - *DateTime: the new DateTime instance (新的 DateTime 实例)
func NewDateTime(t time.Time) *DateTime {
	return &DateTime{Time: t}
}

// Now returns the current DateTime
// Now 返回当前的 DateTime
//
// Returns:
// - *DateTime: the current DateTime (当前的 DateTime)
func Now() *DateTime {
	return NewDateTime(time.Now())
}

// Format formats the DateTime according to the given pattern
// Format 根据给定的模式格式化 DateTime
//
// Parameters:
// - pattern: the format pattern (格式模式)
//
// Returns:
// - string: the formatted DateTime string (格式化后的 DateTime 字符串)
func (dt *DateTime) Format(pattern string) string {
	return dt.Time.Format(pattern)
}

// AddDays adds the given number of days to the DateTime
// AddDays 向 DateTime 添加给定的天数
//
// Parameters:
// - days: the number of days to add (要添加的天数)
//
// Returns:
// - *DateTime: the new DateTime with the added days (添加天数后的新 DateTime)
func (dt *DateTime) AddDays(days int) *DateTime {
	return NewDateTime(dt.Time.AddDate(0, 0, days))
}

// AddMonths adds the given number of months to the DateTime
// AddMonths 向 DateTime 添加给定的月份数
//
// Parameters:
// - months: the number of months to add (要添加的月份数)
//
// Returns:
// - *DateTime: the new DateTime with the added months (添加月份后的新 DateTime)
func (dt *DateTime) AddMonths(months int) *DateTime {
	return NewDateTime(dt.Time.AddDate(0, months, 0))
}

// AddYears adds the given number of years to the DateTime
// AddYears 向 DateTime 添加给定的年份数
//
// Parameters:
// - years: the number of years to add (要添加的年份数)
//
// Returns:
// - *DateTime: the new DateTime with the added years (添加年份后的新 DateTime)
func (dt *DateTime) AddYears(years int) *DateTime {
	return NewDateTime(dt.Time.AddDate(years, 0, 0))
}

// Between calculates the duration between two DateTime instances
// Between 计算两个 DateTime 实例之间的持续时间
//
// Parameters:
// - other: the other DateTime instance to compare with (要比较的另一个 DateTime 实例)
//
// Returns:
// - time.Duration: the duration between the two DateTime instances (两个 DateTime 实例之间的持续时间)
func (dt *DateTime) Between(other *DateTime) time.Duration {
	return dt.Time.Sub(other.Time)
}
