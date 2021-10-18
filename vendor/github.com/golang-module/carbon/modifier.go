package carbon

import (
	"time"
)

// StartOfCentury 本世纪开始时间
func (c Carbon) StartOfCentury() Carbon {
	c.Time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfCentury 本世纪结束时间
func (c Carbon) EndOfCentury() Carbon {
	c.Time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury+999, 12, 31, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfYear 本年开始时间
func (c Carbon) StartOfYear() Carbon {
	c.Time = time.Date(c.Year(), 1, 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfYear 本年结束时间
func (c Carbon) EndOfYear() Carbon {
	c.Time = time.Date(c.Year(), 12, 31, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfQuarter 本季度开始时间
func (c Carbon) StartOfQuarter() Carbon {
	c.Time = time.Date(c.Year(), time.Month(3*c.Quarter()-2), 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfQuarter 本季度结束时间
func (c Carbon) EndOfQuarter() Carbon {
	quarter, day := c.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	c.Time = time.Date(c.Year(), time.Month(3*quarter), day, 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfMonth 本月开始时间
func (c Carbon) StartOfMonth() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfMonth 本月结束时间
func (c Carbon) EndOfMonth() Carbon {
	t := time.Date(c.Year(), time.Month(c.Month()), 1, 23, 59, 59, 999999999, c.Loc)
	c.Time = t.AddDate(0, 1, -1)
	return c
}

// StartOfWeek 本周开始时间
func (c Carbon) StartOfWeek(weekStartDay time.Weekday) Carbon {
	weekDay := c.Time.In(c.Loc).Weekday()
	if weekDay == weekStartDay {
		return c.StartOfDay()
	}
	days := int(weekDay) - int(weekStartDay)
	if weekDay == time.Sunday {
		days = 6
	}
	return c.SubDays(days).StartOfDay()
}

// EndOfWeek 本周结束时间
func (c Carbon) EndOfWeek(weekStartDay time.Weekday) Carbon {
	weekDay := c.Time.In(c.Loc).Weekday()
	if weekStartDay == 0 && weekDay == 6 {
		return c.EndOfDay()
	}
	if weekStartDay == 1 && weekDay == 0 {
		return c.EndOfDay()
	}
	days := 6 - int(weekDay) + int(weekStartDay)
	if weekDay == time.Sunday {
		days = 6
	}
	return c.AddDays(days).EndOfDay()
}

// StartOfDay 本日开始时间
func (c Carbon) StartOfDay() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfDay 本日结束时间
func (c Carbon) EndOfDay() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 23, 59, 59, 999999999, c.Loc)
	return c
}

// StartOfHour 小时开始时间
func (c Carbon) StartOfHour() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 0, 0, 0, c.Loc)
	return c
}

// EndOfHour 小时结束时间
func (c Carbon) EndOfHour() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 59, 59, 999999999, c.Loc)
	return c
}

// StartOfMinute 分钟开始时间
func (c Carbon) StartOfMinute() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 0, 0, c.Loc)
	return c
}

// EndOfMinute 分钟结束时间
func (c Carbon) EndOfMinute() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 59, 999999999, c.Loc)
	return c
}

// StartOfSecond 秒开始时间
func (c Carbon) StartOfSecond() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 0, c.Loc)
	return c
}

// EndOfSecond 秒结束时间
func (c Carbon) EndOfSecond() Carbon {
	c.Time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 999999999, c.Loc)
	return c
}
