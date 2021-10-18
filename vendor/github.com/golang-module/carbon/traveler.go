package carbon

import "time"

// AddDurations 按照持续时间字符串增加时间
// 支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c Carbon) AddDuration(duration string) Carbon {
	if c.Error != nil {
		return c
	}
	td, err := parseByDuration(duration)
	c.Time = c.Time.Add(td)
	c.Error = err
	return c
}

// SubDurations 按照持续时间字符串减少时间
// 支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c Carbon) SubDuration(duration string) Carbon {
	return c.AddDuration("-" + duration)
}

// AddCenturies N世纪后
func (c Carbon) AddCenturies(centuries int) Carbon {
	return c.AddYears(YearsPerCentury * centuries)
}

// AddCenturiesNoOverflow N世纪后(月份不溢出)
func (c Carbon) AddCenturiesNoOverflow(centuries int) Carbon {
	return c.AddYearsNoOverflow(centuries * YearsPerCentury)
}

// AddCentury 1世纪后
func (c Carbon) AddCentury() Carbon {
	return c.AddCenturies(1)
}

// AddCenturyNoOverflow 1世纪后(月份不溢出)
func (c Carbon) AddCenturyNoOverflow() Carbon {
	return c.AddCenturiesNoOverflow(1)
}

// SubCenturies N世纪前
func (c Carbon) SubCenturies(centuries int) Carbon {
	return c.SubYears(YearsPerCentury * centuries)
}

// SubCenturiesNoOverflow N世纪前(月份不溢出)
func (c Carbon) SubCenturiesNoOverflow(centuries int) Carbon {
	return c.SubYearsNoOverflow(centuries * YearsPerCentury)
}

// SubCentury 1世纪前
func (c Carbon) SubCentury() Carbon {
	return c.SubCenturies(1)
}

// SubCenturyNoOverflow 1世纪前(月份不溢出)
func (c Carbon) SubCenturyNoOverflow() Carbon {
	return c.SubCenturiesNoOverflow(1)
}

// AddYears N年后
func (c Carbon) AddYears(years int) Carbon {
	c.Time = c.Time.AddDate(years, 0, 0)
	return c
}

// AddYearsNoOverflow N年后(月份不溢出)
func (c Carbon) AddYearsNoOverflow(years int) Carbon {
	// 获取N年后本月的最后一天
	last := time.Date(c.Year()+years, time.Month(c.Month()), 1, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc).AddDate(0, 1, -1)
	day := c.Day()
	if c.Day() > last.Day() {
		day = last.Day()
	}
	c.Time = time.Date(last.Year(), last.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// AddYear 1年后
func (c Carbon) AddYear() Carbon {
	return c.AddYears(1)
}

// AddYearNoOverflow 1年后(月份不溢出)
func (c Carbon) AddYearNoOverflow() Carbon {
	return c.AddYearsNoOverflow(1)
}

// SubYears N年前
func (c Carbon) SubYears(years int) Carbon {
	return c.AddYears(-years)
}

// SubYearsNoOverflow N年前(月份不溢出)
func (c Carbon) SubYearsNoOverflow(years int) Carbon {
	return c.AddYearsNoOverflow(-years)
}

// SubYear 1年前
func (c Carbon) SubYear() Carbon {
	return c.SubYears(1)
}

// SubYearNoOverflow 1年前(月份不溢出)
func (c Carbon) SubYearNoOverflow() Carbon {
	return c.SubYearsNoOverflow(1)
}

// AddQuarters N季度后
func (c Carbon) AddQuarters(quarters int) Carbon {
	return c.AddMonths(quarters * MonthsPerQuarter)
}

// AddQuartersNoOverflow N季度后(月份不溢出)
func (c Carbon) AddQuartersNoOverflow(quarters int) Carbon {
	return c.AddMonthsNoOverflow(quarters * MonthsPerQuarter)
}

// AddQuarter 1季度后
func (c Carbon) AddQuarter() Carbon {
	return c.AddQuarters(1)
}

// NextQuarters 1季度后(月份不溢出)
func (c Carbon) AddQuarterNoOverflow() Carbon {
	return c.AddQuartersNoOverflow(1)
}

// SubQuarters N季度前
func (c Carbon) SubQuarters(quarters int) Carbon {
	return c.AddQuarters(-quarters)
}

// SubQuartersNoOverflow N季度前(月份不溢出)
func (c Carbon) SubQuartersNoOverflow(quarters int) Carbon {
	return c.AddMonthsNoOverflow(-quarters * MonthsPerQuarter)
}

// SubQuarter 1季度前
func (c Carbon) SubQuarter() Carbon {
	return c.SubQuarters(1)
}

// SubQuarterNoOverflow 1季度前(月份不溢出)
func (c Carbon) SubQuarterNoOverflow() Carbon {
	return c.SubQuartersNoOverflow(1)
}

// AddMonths N月后
func (c Carbon) AddMonths(months int) Carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

// AddMonthsNoOverflow N月后(月份不溢出)
func (c Carbon) AddMonthsNoOverflow(months int) Carbon {
	month := c.Time.Month() + time.Month(months)
	// 获取N月后的最后一天
	last := time.Date(c.Year(), month, 1, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc).AddDate(0, 1, -1)
	day := c.Day()
	if c.Day() > last.Day() {
		day = last.Day()
	}
	c.Time = time.Date(last.Year(), last.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Loc)
	return c
}

// AddMonth 1月后
func (c Carbon) AddMonth() Carbon {
	return c.AddMonths(1)
}

// AddMonthNoOverflow 1月后(月份不溢出)
func (c Carbon) AddMonthNoOverflow() Carbon {
	return c.AddMonthsNoOverflow(1)
}

// SubMonths N月前
func (c Carbon) SubMonths(months int) Carbon {
	return c.AddMonths(-months)
}

// SubMonthsNoOverflow N月前(月份不溢出)
func (c Carbon) SubMonthsNoOverflow(months int) Carbon {
	return c.AddMonthsNoOverflow(-months)
}

// SubMonth 1月前
func (c Carbon) SubMonth() Carbon {
	return c.SubMonths(1)
}

// SubMonthNoOverflow 1月前(月份不溢出)
func (c Carbon) SubMonthNoOverflow() Carbon {
	return c.SubMonthsNoOverflow(1)
}

// AddWeeks N周后
func (c Carbon) AddWeeks(weeks int) Carbon {
	return c.AddDays(weeks * DaysPerWeek)
}

// AddWeek 1周后
func (c Carbon) AddWeek() Carbon {
	return c.AddWeeks(1)
}

// SubWeeks N周前
func (c Carbon) SubWeeks(weeks int) Carbon {
	return c.SubDays(weeks * DaysPerWeek)
}

// SubWeek 1周前
func (c Carbon) SubWeek() Carbon {
	return c.SubWeeks(1)
}

// AddDays N天后
func (c Carbon) AddDays(days int) Carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

// AddDay 1天后
func (c Carbon) AddDay() Carbon {
	return c.AddDays(1)
}

// SubDays N天前
func (c Carbon) SubDays(days int) Carbon {
	return c.AddDays(-days)
}

// SubDay 1天前
func (c Carbon) SubDay() Carbon {
	return c.SubDays(1)
}

// AddHours N小时后
func (c Carbon) AddHours(hours int) Carbon {
	td := time.Duration(hours) * time.Hour
	c.Time = c.Time.Add(td)
	return c
}

// AddHour 1小时后
func (c Carbon) AddHour() Carbon {
	return c.AddHours(1)
}

// SubHours N小时前
func (c Carbon) SubHours(hours int) Carbon {
	return c.AddHours(-hours)
}

// SubHour 1小时前
func (c Carbon) SubHour() Carbon {
	return c.SubHours(1)
}

// AddMinutes N分钟后
func (c Carbon) AddMinutes(minutes int) Carbon {
	td := time.Duration(minutes) * time.Minute
	c.Time = c.Time.Add(td)
	return c
}

// AddMinute 1分钟后
func (c Carbon) AddMinute() Carbon {
	return c.AddMinutes(1)
}

// SubMinutes N分钟前
func (c Carbon) SubMinutes(minutes int) Carbon {
	return c.AddMinutes(-minutes)
}

// SubMinute 1分钟前
func (c Carbon) SubMinute() Carbon {
	return c.SubMinutes(1)
}

// AddSeconds N秒钟后
func (c Carbon) AddSeconds(seconds int) Carbon {
	td := time.Duration(seconds) * time.Second
	c.Time = c.Time.Add(td)
	return c
}

// AddSecond 1秒钟后
func (c Carbon) AddSecond() Carbon {
	return c.AddSeconds(1)
}

// SubSeconds N秒钟前
func (c Carbon) SubSeconds(seconds int) Carbon {
	return c.AddSeconds(-seconds)
}

// SubSecond 1秒钟前
func (c Carbon) SubSecond() Carbon {
	return c.SubSeconds(1)
}
