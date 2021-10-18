package carbon

// DaysInYear 获取本年的总天数
func (c Carbon) DaysInYear() int {
	if c.IsZero() {
		return 0
	}
	days := DaysPerNormalYear
	if c.IsLeapYear() {
		days = DaysPerLeapYear
	}
	return days
}

// DaysInMonth 获取本月的总天数
func (c Carbon) DaysInMonth() int {
	if c.IsZero() {
		return 0
	}
	return c.EndOfMonth().Time.In(c.Loc).Day()
}

// MonthOfYear 获取本年的第几月(从1开始)
func (c Carbon) MonthOfYear() int {
	if c.IsZero() {
		return 0
	}
	return int(c.Time.In(c.Loc).Month())
}

// DayOfYear 获取本年的第几天(从1开始)
func (c Carbon) DayOfYear() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).YearDay()
}

// DayOfMonth 获取本月的第几天(从1开始)
func (c Carbon) DayOfMonth() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Day()
}

// DayOfWeek 获取本周的第几天(从1开始)
func (c Carbon) DayOfWeek() int {
	if c.IsZero() {
		return 0
	}
	day := int(c.Time.In(c.Loc).Weekday())
	if day == 0 {
		return DaysPerWeek
	}
	return day
}

// WeekOfYear 获取本年的第几周(从1开始)
func (c Carbon) WeekOfYear() int {
	if c.IsZero() {
		return 0
	}
	_, week := c.Time.In(c.Loc).ISOWeek()
	return week
}

// WeekOfMonth 获取本月的第几周(从1开始)
func (c Carbon) WeekOfMonth() int {
	if c.IsZero() {
		return 0
	}
	day := c.Time.In(c.Loc).Day()
	if day < DaysPerWeek {
		return 1
	}
	return day%DaysPerWeek + 1
}

// Century 获取当前世纪
func (c Carbon) Century() int {
	if c.IsZero() {
		return 0
	}
	return c.Year()/YearsPerCentury + 1
}

// Year 获取当前年
func (c Carbon) Year() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Year()
}

// Quarter 获取当前季度
func (c Carbon) Quarter() int {
	if c.IsZero() {
		return 0
	}
	switch {
	case c.Month() >= 10:
		return 4
	case c.Month() >= 7:
		return 3
	case c.Month() >= 4:
		return 2
	case c.Month() >= 1:
		return 1
	default:
		return 0
	}
}

// Month 获取当前月
func (c Carbon) Month() int {
	if c.IsZero() {
		return 0
	}
	return c.MonthOfYear()
}

// Week 获取当前周(从0开始)
func (c Carbon) Week() int {
	if c.IsZero() {
		return -1
	}
	return int(c.Time.In(c.Loc).Weekday())
}

// Day 获取当前日
func (c Carbon) Day() int {
	if c.IsZero() {
		return 0
	}
	return c.DayOfMonth()
}

// Hour 获取当前小时
func (c Carbon) Hour() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Hour()
}

// Minute 获取当前分钟数
func (c Carbon) Minute() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Minute()
}

// Second 获取当前秒数
func (c Carbon) Second() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Second()
}

// Millisecond 获取当前毫秒数
func (c Carbon) Millisecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e6
}

// Microsecond 获取当前微秒数
func (c Carbon) Microsecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond() / 1e9
}

// Nanosecond 获取当前纳秒数
func (c Carbon) Nanosecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.In(c.Loc).Nanosecond()
}

// Timezone 获取时区
func (c Carbon) Timezone() string {
	return c.Loc.String()
}

// Locale 获取语言区域
func (c Carbon) Locale() string {
	return c.Lang.locale
}

// Age 获取年龄
func (c Carbon) Age() int {
	if c.IsZero() {
		return 0
	}
	now := Now()
	if c.ToTimestamp() > now.ToTimestamp() {
		return 0
	}
	return int(c.DiffInYears(now))
}
