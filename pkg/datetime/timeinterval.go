// Package datetime provides utility functions for handling time intervals and durations.
// This package simplifies measuring the execution time of code.
//
// Package datetime 提供了处理时间间隔和持续时间的实用函数。
// 该包简化了测量代码执行时间的过程。
package datetime

import (
	"time"
)

// TimeInterval is a simple timer class used to calculate the duration of code execution
//
// TimeInterval 是一个简单的计时器类，用于计算代码执行的持续时间
type TimeInterval struct {
	start time.Time
}

// NewTimeInterval creates a new TimeInterval instance and starts the timer
//
// # NewTimeInterval 创建一个新的 TimeInterval 实例并启动计时器
//
// Returns:
// - *TimeInterval: a new TimeInterval instance (一个新的 TimeInterval 实例)
func NewTimeInterval() *TimeInterval {
	return &TimeInterval{start: time.Now()}
}

// Elapsed returns the elapsed time in the specified unit
//
// # Elapsed 返回指定单位的经过时间
//
// Parameters:
// - unit: the unit of time to return (返回的时间单位)
//
// Returns:
// - int64: the elapsed time in the specified unit (指定单位的经过时间)
func (ti *TimeInterval) Elapsed(unit DateUnit) int64 {
	duration := time.Since(ti.start)
	return duration.Nanoseconds() / (unit.GetMillis() * 1e6)
}

// ElapsedMillis returns the elapsed time in milliseconds
//
// # ElapsedMillis 返回以毫秒为单位的经过时间
//
// Returns:
// - int64: the elapsed time in milliseconds (以毫秒为单位的经过时间)
func (ti *TimeInterval) ElapsedMillis() int64 {
	return ti.Elapsed(Millisecond)
}

// ElapsedSeconds returns the elapsed time in seconds
//
// # ElapsedSeconds 返回以秒为单位的经过时间
//
// Returns:
// - int64: the elapsed time in seconds (以秒为单位的经过时间)
func (ti *TimeInterval) ElapsedSeconds() int64 {
	return ti.Elapsed(Second)
}

// ElapsedMinutes returns the elapsed time in minutes
//
// # ElapsedMinutes 返回以分钟为单位的经过时间
//
// Returns:
// - int64: the elapsed time in minutes (以分钟为单位的经过时间)
func (ti *TimeInterval) ElapsedMinutes() int64 {
	return ti.Elapsed(Minute)
}

// ElapsedHours returns the elapsed time in hours
//
// # ElapsedHours 返回以小时为单位的经过时间
//
// Returns:
// - int64: the elapsed time in hours (以小时为单位的经过时间)
func (ti *TimeInterval) ElapsedHours() int64 {
	return ti.Elapsed(Hour)
}

// ElapsedDays returns the elapsed time in days
//
// # ElapsedDays 返回以天为单位的经过时间
//
// Returns:
// - int64: the elapsed time in days (以天为单位的经过时间)
func (ti *TimeInterval) ElapsedDays() int64 {
	return ti.Elapsed(Day)
}

// ElapsedWeeks returns the elapsed time in weeks
//
// # ElapsedWeeks 返回以周为单位的经过时间
//
// Returns:
// - int64: the elapsed time in weeks (以周为单位的经过时间)
func (ti *TimeInterval) ElapsedWeeks() int64 {
	return ti.Elapsed(Week)
}
