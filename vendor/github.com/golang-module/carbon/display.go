package carbon

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

// ToTimestamp ToTimestampWithSecond的简称
func (c Carbon) ToTimestamp() int64 {
	return c.ToTimestampWithSecond()
}

// ToTimestampWithSecond 输出秒级时间戳
func (c Carbon) ToTimestampWithSecond() int64 {
	return c.Time.Unix()
}

// ToTimestampWithMillisecond 输出毫秒级时间戳
func (c Carbon) ToTimestampWithMillisecond() int64 {
	return c.Time.UnixNano() / int64(time.Millisecond)
}

// ToTimestampWithMicrosecond 输出微秒级时间戳
func (c Carbon) ToTimestampWithMicrosecond() int64 {
	return c.Time.UnixNano() / int64(time.Microsecond)
}

// ToTimestampWithNanosecond 输出纳秒级时间戳
func (c Carbon) ToTimestampWithNanosecond() int64 {
	return c.Time.UnixNano()
}

// String 实现 Stringer 接口
func (c Carbon) String() string {
	if c.IsZero() {
		return ""
	}
	return c.ToDateTimeString()
}

// ToString 输出"2006-01-02 15:04:05.999999999 -0700 MST"格式字符串
func (c Carbon) ToString() string {
	return c.Time.In(c.Loc).String()
}

// ToUtcString 输出0时区时间字符串
func (c Carbon) ToUtcString() string {
	return c.Time.In(c.Loc).UTC().String()
}

// ToMonthString 输出完整月份字符串，支持i18n
func (c Carbon) ToMonthString() string {
	if len(c.Lang.resources) == 0 && c.Lang.SetLocale(defaultLocale) != nil {
		return ""
	}
	if months, ok := c.Lang.resources["months"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Month()-1]
	}
	return ""
}

// ToShortMonthString 输出缩写月份字符串，支持i18n
func (c Carbon) ToShortMonthString() string {
	if len(c.Lang.resources) == 0 && c.Lang.SetLocale(defaultLocale) != nil {
		return ""
	}
	if months, ok := c.Lang.resources["months_short"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Month()-1]
	}
	if months, ok := c.Lang.resources["months"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Month()-1]
	}
	return ""
}

// ToWeekString 输出完整星期字符串，支持i18n
func (c Carbon) ToWeekString() string {
	if len(c.Lang.resources) == 0 && c.Lang.SetLocale(defaultLocale) != nil {
		return ""
	}
	if months, ok := c.Lang.resources["weeks"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Week()]
	}
	return ""
}

// ToShortWeekString 输出缩写星期字符串，支持i18n
func (c Carbon) ToShortWeekString() string {
	if len(c.Lang.resources) == 0 && c.Lang.SetLocale(defaultLocale) != nil {
		return ""
	}
	if months, ok := c.Lang.resources["weeks_short"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Week()]
	}
	if months, ok := c.Lang.resources["weeks"]; ok {
		slice := strings.Split(months, "|")
		return slice[c.Week()]
	}
	return ""
}

// Format ToFormatString的简称
func (c Carbon) Format(format string) string {
	return c.ToFormatString(format)
}

// ToFormatString 输出指定格式时间
func (c Carbon) ToFormatString(format string) string {
	if c.IsZero() {
		return ""
	}
	runes := []rune(format)
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(runes); i++ {
		if layout, ok := formats[byte(runes[i])]; ok {
			buffer.WriteString(c.Time.In(c.Loc).Format(layout))
		} else {
			switch runes[i] {
			case '\\': // 原样输出，不解析
				buffer.WriteRune(runes[i+1])
				i += 2
				continue
			case 'W': // ISO-8601 格式数字表示的年份中的第几周，取值范围 1-52
				buffer.WriteString(strconv.Itoa(c.WeekOfYear()))
			case 'N': // ISO-8601 格式数字表示的星期中的第几天，取值范围 1-7
				buffer.WriteString(strconv.Itoa(c.DayOfWeek()))
			case 'S': // 月份中第几天的英文缩写后缀，如st, nd, rd, th
				suffix := "th"
				switch c.Day() {
				case 1, 21, 31:
					suffix = "st"
				case 2, 22:
					suffix = "nd"
				case 3, 23:
					suffix = "rd"
				}
				buffer.WriteString(suffix)
			case 'L': // 是否为闰年，如果是闰年为 1，否则为 0
				if c.IsLeapYear() {
					buffer.WriteString("1")
				} else {
					buffer.WriteString("0")
				}
			case 'G': // 数字表示的小时，24 小时格式，没有前导零，取值范围 0-23
				buffer.WriteString(strconv.Itoa(c.Hour()))
			case 'U': // 秒级时间戳，如 1611818268
				buffer.WriteString(strconv.FormatInt(c.ToTimestamp(), 10))
			case 'u': // 数字表示的毫秒，如 999
				buffer.WriteString(strconv.Itoa(c.Millisecond()))
			case 'w': // 数字表示的星期中的第几天，取值范围 0-6
				buffer.WriteString(strconv.Itoa(c.DayOfWeek() - 1))
			case 't': // 指定的月份有几天，取值范围 28-31
				buffer.WriteString(strconv.Itoa(c.DaysInMonth()))
			case 'z': // 年份中的第几天，取值范围 0-365
				buffer.WriteString(strconv.Itoa(c.DayOfYear() - 1))
			case 'e': // 当前时区，如 UTC，GMT，Atlantic/Azores
				buffer.WriteString(c.Timezone())
			default:
				buffer.WriteRune(runes[i])
			}
		}
	}
	return buffer.String()
}

// ToDayDateTimeString 输出天数日期时间字符串
func (c Carbon) ToDayDateTimeString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(DayDateTimeFormat)
}

// ToDayDateTimeStringWithTimezone 输出指定时区的天数日期时间字符串
func (c Carbon) ToDayDateTimeStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(DayDateTimeFormat)
}

// ToDateTimeString 输出日期时间字符串
func (c Carbon) ToDateTimeString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(DateTimeFormat)
}

// ToDateTimeStringWithTimezone 输出指定时区的日期时间字符串
func (c Carbon) ToDateTimeStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(DateTimeFormat)
}

// ToDateString 输出日期字符串
func (c Carbon) ToDateString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(DateFormat)
}

// ToDateStringWithTimezone 输出指定时区的日期字符串
func (c Carbon) ToDateStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(DateFormat)
}

// ToTimeString 输出时间字符串
func (c Carbon) ToTimeString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(TimeFormat)
}

// ToTimeStringWithTimezone 输出指定时区的时间字符串
func (c Carbon) ToTimeStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(TimeFormat)
}

// ToAtomString 输出Atom格式字符串
func (c Carbon) ToAtomString() string {
	return c.ToRfc3339String()
}

// ToAtomStringWithTimezone 输出指定时区的Atom格式字符串
func (c Carbon) ToAtomStringWithTimezone(timezone string) string {
	return c.ToRfc3339StringWithTimezone(timezone)
}

// ToAnsicString 输出Ansic格式字符串
func (c Carbon) ToAnsicString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(AnsicFormat)
}

// ToAnsicStringWithTimezone 输出指定时区的Ansic格式字符串
func (c Carbon) ToAnsicStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(AnsicFormat)
}

// ToCookieString 输出Cookie格式字符串
func (c Carbon) ToCookieString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(CookieFormat)
}

// ToCookieStringWithTimezone 输出指定时区的Cookie格式字符串
func (c Carbon) ToCookieStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(CookieFormat)
}

// ToRssString 输出RSS格式字符串
func (c Carbon) ToRssString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RssFormat)
}

// ToRssStringWithTimezone 输出指定时区的RSS格式字符串
func (c Carbon) ToRssStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RssFormat)
}

// ToW3cString 输出W3C格式字符串
func (c Carbon) ToW3cString() string {
	return c.ToRfc3339String()
}

// ToW3cStringWithTimezone 输出指定时区的W3C格式字符串
func (c Carbon) ToW3cStringWithTimezone(timezone string) string {
	return c.ToRfc3339StringWithTimezone(timezone)
}

// ToUnixDateString 输出UnixDate格式字符串
func (c Carbon) ToUnixDateString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(UnixDateFormat)
}

// ToUnixDateStringWithTimezone 输出指定时区的UnixDate格式字符串
func (c Carbon) ToUnixDateStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(UnixDateFormat)
}

// ToUnixDateString 输出RubyDate格式字符串
func (c Carbon) ToRubyDateString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RubyDateFormat)
}

// ToUnixDateStringWithTimezone 输出指定时区的RubyDate格式字符串
func (c Carbon) ToRubyDateStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RubyDateFormat)
}

// ToKitchenString 输出Kitchen格式字符串
func (c Carbon) ToKitchenString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(KitchenFormat)
}

// ToKitchenStringWithTimezone 输出指定时区的Kitchen格式字符串
func (c Carbon) ToKitchenStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(KitchenFormat)
}

// ToRfc822String 输出RFC822格式字符串
func (c Carbon) ToRfc822String() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC822Format)
}

// ToRfc822StringWithTimezone 输出指定时区的RFC822格式字符串
func (c Carbon) ToRfc822StringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC822Format)
}

// ToRfc822zString 输出RFC822Z格式字符串
func (c Carbon) ToRfc822zString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC822ZFormat)
}

// ToRfc822zStringWithTimezone 输出指定时区的RFC822Z格式字符串
func (c Carbon) ToRfc822zStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC822ZFormat)
}

// ToRfc850String 输出RFC850格式字符串
func (c Carbon) ToRfc850String() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC850Format)
}

// ToRfc850StringWithTimezone 输出指定时区的RFC850格式字符串
func (c Carbon) ToRfc850StringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC850Format)
}

// ToRfc1036String 输出RFC1036格式字符串
func (c Carbon) ToRfc1036String() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC1036Format)
}

// ToRfc1036StringWithTimezone 输出指定时区的RFC1036格式字符串
func (c Carbon) ToRfc1036StringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC1036Format)
}

// ToRfc1123String 输出RFC1123格式字符串
func (c Carbon) ToRfc1123String() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC1123Format)
}

// ToRfc1123StringWithTimezone 输出指定时区的RFC1123格式字符串
func (c Carbon) ToRfc1123StringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC1123Format)
}

// ToRfc1123ZString 输出RFC1123Z格式字符串
func (c Carbon) ToRfc1123ZString() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC1123ZFormat)
}

// ToRfc1123ZStringWithTimezone 输出指定时区的RFC1123Z格式字符串
func (c Carbon) ToRfc1123ZStringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC1123ZFormat)
}

// ToRfc2822String 输出RFC2822格式字符串
func (c Carbon) ToRfc2822String() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC2822Format)
}

// ToRfc2822StringWithTimezone 输出指定时区的RFC2822格式字符串
func (c Carbon) ToRfc2822StringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC2822Format)
}

// ToRfc3339String 输出RFC3339格式字符串
func (c Carbon) ToRfc3339String() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC3339Format)
}

// ToRfc3339StringWithTimezone 输出指定时区的RFC3339格式字符串
func (c Carbon) ToRfc3339StringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC3339Format)
}

// ToRfc7231String 输出RFC7231格式字符串
func (c Carbon) ToRfc7231String() string {
	if c.IsZero() {
		return ""
	}
	return c.Time.In(c.Loc).Format(RFC7231Format)
}

// ToRfc7231StringWithTimezone 输出指定时区的RFC7231格式字符串
func (c Carbon) ToRfc7231StringWithTimezone(timezone string) string {
	if c.IsZero() {
		return ""
	}
	loc, _ := time.LoadLocation(timezone)
	return c.Time.In(c.Loc).In(loc).Format(RFC7231Format)
}
