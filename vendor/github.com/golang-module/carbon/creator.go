package carbon

import (
	"strconv"
	"time"
)

// CreateFromTimestamp 从时间戳创建 Carbon 实例
func (c Carbon) CreateFromTimestamp(timestamp int64) Carbon {
	ts := timestamp
	switch len(strconv.FormatInt(timestamp, 10)) {
	case 10:
		ts = timestamp
	case 13:
		ts = timestamp / 1e3
	case 16:
		ts = timestamp / 1e6
	case 19:
		ts = timestamp / 1e9
	default:
		ts = 0
	}
	c.Time = time.Unix(ts, 0)
	return c
}

// CreateFromTimestamp 从时间戳创建 Carbon 实例(默认时区)
func CreateFromTimestamp(timestamp int64) Carbon {
	return New().CreateFromTimestamp(timestamp)
}

// CreateFromDateTime 从年月日时分秒创建 Carbon 实例
func (c Carbon) CreateFromDateTime(year int, month int, day int, hour int, minute int, second int) Carbon {
	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, 0, c.Loc)
	return c
}

// CreateFromDateTime 从年月日时分秒创建 Carbon 实例(默认时区)
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int) Carbon {
	return New().CreateFromDateTime(year, month, day, hour, minute, second)
}

// CreateFromDate 从年月日创建 Carbon 实例
func (c Carbon) CreateFromDate(year int, month int, day int) Carbon {
	hour, minute, second := time.Now().Clock()
	c.Time = time.Date(year, time.Month(month), day, hour, minute, second, 0, c.Loc)
	return c
}

// CreateFromDate 从年月日创建 Carbon 实例(默认时区)
func CreateFromDate(year int, month int, day int) Carbon {
	return New().CreateFromDate(year, month, day)
}

// CreateFromTime 从时分秒创建 Carbon 实例
func (c Carbon) CreateFromTime(hour int, minute int, second int) Carbon {
	year, month, day := time.Now().Date()
	c.Time = time.Date(year, month, day, hour, minute, second, 0, c.Loc)
	return c
}

// CreateFromTime 从时分秒创建 Carbon 实例(默认时区)
func CreateFromTime(hour int, minute int, second int) Carbon {
	return New().CreateFromTime(hour, minute, second)
}
