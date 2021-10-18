package carbon

import "time"

// IsZero 是否是零值
func (c Carbon) IsZero() bool {
	return c.Time.IsZero()
}

// IsNow 是否是当前时间
func (c Carbon) IsNow() bool {
	return c.ToTimestamp() == c.Now().ToTimestamp()
}

// IsFuture 是否是未来时间
func (c Carbon) IsFuture() bool {
	return c.ToTimestamp() > c.Now().ToTimestamp()
}

// IsPast 是否是过去时间
func (c Carbon) IsPast() bool {
	return c.ToTimestamp() < c.Now().ToTimestamp()
}

// IsLeapYear 是否是闰年
func (c Carbon) IsLeapYear() bool {
	year := c.Time.In(c.Loc).Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsLongYear 是否是长年
func (c Carbon) IsLongYear() bool {
	_, w := time.Date(c.Year(), time.December, 31, 0, 0, 0, 0, c.Loc).ISOWeek()
	return w == weeksPerLongYear
}

// IsJanuary 是否是一月
func (c Carbon) IsJanuary() bool {
	return c.Time.In(c.Loc).Month() == time.January
}

// IsMonday 是否是二月
func (c Carbon) IsFebruary() bool {
	return c.Time.In(c.Loc).Month() == time.February
}

// IsMarch 是否是三月
func (c Carbon) IsMarch() bool {
	return c.Time.In(c.Loc).Month() == time.March
}

// IsApril 是否是四月
func (c Carbon) IsApril() bool {
	return c.Time.In(c.Loc).Month() == time.April
}

// IsMay 是否是五月
func (c Carbon) IsMay() bool {
	return c.Time.In(c.Loc).Month() == time.May
}

// IsJune 是否是六月
func (c Carbon) IsJune() bool {
	return c.Time.In(c.Loc).Month() == time.June
}

// IsJuly 是否是七月
func (c Carbon) IsJuly() bool {
	return c.Time.In(c.Loc).Month() == time.July
}

// IsAugust 是否是八月
func (c Carbon) IsAugust() bool {
	return c.Time.In(c.Loc).Month() == time.August
}

// IsSeptember 是否是九月
func (c Carbon) IsSeptember() bool {
	return c.Time.In(c.Loc).Month() == time.September
}

// IsOctober 是否是十月
func (c Carbon) IsOctober() bool {
	return c.Time.In(c.Loc).Month() == time.October
}

// IsNovember 是否是十一月
func (c Carbon) IsNovember() bool {
	return c.Time.In(c.Loc).Month() == time.November
}

// IsDecember 是否是十二月
func (c Carbon) IsDecember() bool {
	return c.Time.In(c.Loc).Month() == time.December
}

// IsMonday 是否是周一
func (c Carbon) IsMonday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Monday
}

// IsTuesday 是否是周二
func (c Carbon) IsTuesday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Tuesday
}

// IsWednesday 是否是周三
func (c Carbon) IsWednesday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Wednesday
}

// IsThursday 是否是周四
func (c Carbon) IsThursday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Thursday
}

// IsFriday 是否是周五
func (c Carbon) IsFriday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Friday
}

// IsSaturday 是否是周六
func (c Carbon) IsSaturday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Saturday
}

// IsSunday 是否是周日
func (c Carbon) IsSunday() bool {
	return c.Time.In(c.Loc).Weekday() == time.Sunday
}

// IsWeekday 是否是工作日
func (c Carbon) IsWeekday() bool {
	return !c.IsSaturday() && !c.IsSunday()
}

// IsWeekend 是否是周末
func (c Carbon) IsWeekend() bool {
	return c.IsSaturday() || c.IsSunday()
}

// IsYesterday 是否是昨天
func (c Carbon) IsYesterday() bool {
	return c.ToDateString() == Now().SubDay().ToDateString()
}

// IsToday 是否是今天
func (c Carbon) IsToday() bool {
	return c.ToDateString() == c.Now().ToDateString()
}

// IsTomorrow 是否是明天
func (c Carbon) IsTomorrow() bool {
	return c.ToDateString() == Now().AddDay().ToDateString()
}

// Compare 时间比较
func (c Carbon) Compare(operator string, t Carbon) bool {
	switch operator {
	case "=":
		return c.Eq(t)
	case "<>":
		return !c.Eq(t)
	case "!=":
		return !c.Eq(t)
	case ">":
		return c.Gt(t)
	case ">=":
		return c.Gte(t)
	case "<":
		return c.Lt(t)
	case "<=":
		return c.Lte(t)
	}
	return false
}

// Gt 大于
func (c Carbon) Gt(t Carbon) bool {
	return c.Time.After(t.Time)
}

// Lt 小于
func (c Carbon) Lt(t Carbon) bool {
	return c.Time.Before(t.Time)
}

// Eq 等于
func (c Carbon) Eq(t Carbon) bool {
	return c.Time.Equal(t.Time)
}

// Ne 不等于
func (c Carbon) Ne(t Carbon) bool {
	return !c.Eq(t)
}

// Gte 大于等于
func (c Carbon) Gte(t Carbon) bool {
	return c.Gt(t) || c.Eq(t)
}

// Lte 小于等于
func (c Carbon) Lte(t Carbon) bool {
	return c.Lt(t) || c.Eq(t)
}

// Between 是否在两个时间之间(不包括这两个时间)
func (c Carbon) Between(start Carbon, end Carbon) bool {
	if c.Gt(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStartTime 是否在两个时间之间(包括开始时间)
func (c Carbon) BetweenIncludedStartTime(start Carbon, end Carbon) bool {
	if c.Gte(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEndTime 是否在两个时间之间(包括结束时间)
func (c Carbon) BetweenIncludedEndTime(start Carbon, end Carbon) bool {
	if c.Gt(start) && c.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth 是否在两个时间之间(包括这两个时间)
func (c Carbon) BetweenIncludedBoth(start Carbon, end Carbon) bool {
	if c.Gte(start) && c.Lte(end) {
		return true
	}
	return false
}
