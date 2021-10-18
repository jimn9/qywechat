# Carbon  #
中文 | [English](./README.en.md)

carbon 是一个轻量级、语义化、对开发者友好的 Golang 时间处理库，支持链式调用和 gorm、xorm、zorm 等主流 orm。

如果您觉得不错，请给个 star 吧

github:[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

gitee:[gitee.com/go-package/carbon](https://gitee.com/go-package/carbon "gitee.com/go-package/carbon")

#### 安装使用
```go
// 使用 github 库
go get -u github.com/golang-module/carbon

import (
    "github.com/golang-module/carbon"
)

// 使用 gitee 库
go get -u gitee.com/go-package/carbon

import (
    "gitee.com/go-package/carbon"
)
```

#### 用法示例
> 默认时区为 Local，即服务器所在时区，假设当前时间为 2020-08-05 13:14:15

##### 昨天、今天、明天
```go
// 今天此刻
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今天日期
carbon.Now().ToDateString() // 2020-08-05
// 今天时间
carbon.Now().ToTimeString() // 13:14:15
// 今天秒级时间戳
carbon.Now().ToTimestamp() // 1596604455
carbon.Now().ToTimestampWithSecond() // 1596604455
// 今天毫秒级时间戳
carbon.Now().ToTimestampWithMillisecond() // 1596604455000
// 今天微秒级时间戳
carbon.Now().ToTimestampWithMicrosecond() // 1596604455000000
// 今天纳秒级时间戳
carbon.Now().ToTimestampWithNanosecond() // 1596604455000000000
// 指定时区的今天此刻
carbon.SetTimezone(Carbon.NewYork).Now().ToDateTimeString() // 2020-08-05 01:14:15

// 昨天此刻
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨天日期
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨天时间
carbon.Yesterday().ToTimeString() // 13:14:15
// 昨天秒级时间戳
carbon.Yesterday().ToTimestamp() // 1596518055
carbon.Yesterday().ToTimestampWithSecond() // 1596518055
// 明天毫秒级时间戳
carbon.Yesterday().ToTimestampWithMillisecond() // 1596518055000
// 明天微秒级时间戳
carbon.Yesterday().ToTimestampWithMicrosecond() // 1596518055000000
// 明天纳秒级时间戳
carbon.Yesterday().ToTimestampWithNanosecond() // 1596518055000000000
// 指定时区的昨天此刻
carbon.SetTimezone(Carbon.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 01:14:15
// 指定日期的昨天此刻
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15

// 明天此刻
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明天日期
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明天时间
carbon.Tomorrow().ToTimeString() // 13:14:15
// 明天秒级时间戳
carbon.Tomorrow().ToTimestamp() // 1596690855
carbon.Tomorrow().ToTimestampWithSecond() // 1596690855
// 明天毫秒级时间戳
carbon.Tomorrow().ToTimestampWithMillisecond() // 1596690855000
// 明天微秒级时间戳
carbon.Tomorrow().ToTimestampWithMicrosecond() // 1596690855000000
// 明天纳秒级时间戳
carbon.Tomorrow().ToTimestampWithNanosecond() // 1596690855000000000
// 指定时区的明天此刻
carbon.SetTimezone(Carbon.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 01:14:15
// 指定日期的明天此刻
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
```

##### 创建 Carbon 实例
```go
// 从秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
// 从毫秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
// 从微秒级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
// 从纳级时间戳创建 Carbon 实例
carbon.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15

// 从年月日时分秒创建 Carbon 实例
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// 从年月日创建 Carbon 实例(时分秒默认为当前时分秒)
carbon.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)
carbon.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
```

##### 将标准格式时间字符串解析成 Carbon 实例
```go
carbon.Parse("").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空字符串
carbon.Parse("0000-00-00").ToDateTimeString() // 空字符串
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
carbon.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 00:00:00
```

##### 将特殊格式时间字符串解析成 Carbon 实例
```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "It is Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
```

##### 将布局时间字符串解析成 Carbon 实例
```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
```

##### carbon 和 time.Time 互转
```go
// 将 time.Time 转换成 Carbon
carbon.Time2Carbon(time.Now())
// 将 Carbon 转换成 time.Time
carbon.Now().Carbon2Time()
```

##### 开始时间、结束时间
```go
// 本世纪开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// 本世纪结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// 本年开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// 本年结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// 本季度开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// 本季度结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// 本月开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// 本月结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// 本周开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfWeek(time.Sunday).ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").StartOfWeek(time.Monday).ToDateTimeString() // 2020-08-03 00:00:00
// 本周结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfWeek(time.Sunday).ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").EndOfWeek(time.Monday).ToDateTimeString() // 2020-08-09 23:59:59

// 本日开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// 本日结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// 本小时开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// 本小时结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// 本分钟开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// 本分钟结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// 本秒开始时间
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.0
// 本秒结束时间
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.999
```

##### 时间旅行
```go
// 三世纪后
carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// 三世纪后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15

// 一世纪后
carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// 一世纪后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15

// 三世纪前
carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// 三世纪前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15

// 一世纪前
carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// 一世纪前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15

// 三年后
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// 三年后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15

// 一年后
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// 一年后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15

// 三年前
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// 三年前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15

// 一年前
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// 一年前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// 三季度后
carbon.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
// 三季度后(月份不溢出)
carbon.Parse("2019-08-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2019-02-29 13:14:15

// 一季度后
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// 一季度后(月份不溢出)
carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三季度前
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// 三季度前(月份不溢出)
carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15

// 一季度前
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// 一季度前(月份不溢出)
carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三月后
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// 三月后(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15

// 一月后
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一月后(月份不溢出)
carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三月前
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// 三月前(月份不溢出)
carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15

// 一月前
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一月前(月份不溢出)
carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三周后
carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// 一周后
carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15

// 三周前
carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// 一周前
carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// 三天后
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// 一天后
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15

// 三天前
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// 一天前
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// 三小时后
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// 二小时半后
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// 一小时后
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15

// 三小时前
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 二小时半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// 一小时前
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// 三分钟后
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 二分钟半后
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// 一分钟后
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15

// 三分钟前
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 二分钟半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
carbon.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
// 一分钟前
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// 三秒钟后
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// 二秒钟半后
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// 一秒钟后
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16

// 三秒钟前
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// 二秒钟半前
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// 一秒钟前
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14
```

##### 时间差
```go
// 相差多少周
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
// 相差多少周（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInWeeksWithAbs(carbon.Parse("2020-07-28 13:14:15")) // 1

// 相差多少天
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
// 相差多少天（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInDaysWithAbs(carbon.Parse("2020-08-04 13:14:15")) // 1

// 相差多少小时
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
// 相差多少小时（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInHoursWithAbs(carbon.Parse("2020-08-05 12:14:15")) // 1

// 相差多少分
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
// 相差多少分（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInMinutesWithAbs(carbon.Parse("2020-08-05 13:13:15")) // 1

// 相差多少秒
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
// 相差多少秒（绝对值）
carbon.Parse("2020-08-05 13:14:15").DiffInSecondsWithAbs(carbon.Parse("2020-08-05 13:14:14")) // 1

// 对人类友好的可读格式时间差(需要先把lang目录复制到项目目录下)
carbon.Parse("2020-08-05 13:14:15").DiffForHumans()) // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year before
carbon.Parse("2019-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years before
carbon.Parse("2018-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year after
carbon.Parse("2022-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years after
```

##### 时间判断
```go
// 是否是零值时间
carbon.Parse("").IsZero() // true
carbon.Parse("0").IsZero() // true
carbon.Parse("0000-00-00 00:00:00").IsZero() // true
carbon.Parse("0000-00-00").IsZero() // true
carbon.Parse("00:00:00").IsZero() // true
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false

// 是否是当前时间
carbon.Now().IsNow() // true
// 是否是未来时间
carbon.Tomorrow().IsFuture() // true
// 是否是过去时间
carbon.Yesterday().IsPast() // true

// 是否是闰年
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// 是否是长年
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 是否是一月
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 是否是二月
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 是否是三月
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// 是否是四月
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// 是否是五月
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// 是否是六月
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// 是否是七月
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// 是否是八月
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// 是否是九月
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 是否是十月
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// 是否是十一月
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// 是否是十二月
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// 是否是周一
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// 是否是周二
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 是否是周三
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 是否是周四
carbon.Parse("2020-08-05 13:14:15").IsThursday()  // false
// 是否是周五
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// 是否是周六
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 是否是周日
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

// 是否是工作日
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 是否是周末
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 是否是昨天
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// 是否是今天
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// 是否是明天
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// 是否大于
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// 是否小于
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// 是否等于
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// 是否不等于
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// 是否大于等于
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// 是否小于等于
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true

// 是否在两个时间之间(不包括这两个时间)
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 是否在两个时间之间(包括开始时间)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStartTime(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStartTime(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 是否在两个时间之间(包括结束时间)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEndTime(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEndTime(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 是否在两个时间之间(包括这两个时间)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true

```

##### 时间输出
```go
// 输出秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestamp() // 1596604455
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithSecond() // 1596604455
// 输出毫秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMillisecond() // 1596604455000
// 输出微秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithMicrosecond() // 1596604455000000
// 输出纳秒级时间戳
carbon.Parse("2020-08-05 13:14:15").ToTimestampWithNanosecond() // 1596604455000000000

// 输出日期时间字符串
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToDateTimeStringWithTimezone(carbon.LosAngeles) // 2020-08-04 22:14:15
// 输出日期字符串
carbon.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
carbon.Parse("2020-08-05 13:14:15").ToDateStringWithTimezone(carbon.LosAngeles) // 2020-08-04
// 输出时间字符串
carbon.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15
carbon.Parse("2020-08-05 13:14:15").ToTimeStringWithTimezone(carbon.LosAngeles) // 22:14:15

// 输出 Ansic 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
carbon.Parse("2020-08-05 13:14:15").ToAnsicStringWithTimezone(carbon.LosAngeles) // Tue Aug  4 22:14:15 2020
// 输出 Atom 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToAtomStringWithTimezone(carbon.LosAngeles) // 2020-08-04T22:14:15-07:00
// 输出 UnixDate 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
carbon.Parse("2020-08-05 13:14:15").ToUnixDateStringWithTimezone(carbon.LosAngeles) // Tue Aug  4 22:14:15 PDT 2020
// 输出 RubyDate 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
carbon.Parse("2020-08-05 13:14:15").ToRubyDateStringWithTimezone(carbon.LosAngeles) // Tue Aug 04 22:14:15 -0700 2020
// 输出 Kitchen 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
carbon.Parse("2020-08-05 13:14:15").ToKitchenStringWithTimezone(carbon.LosAngeles) // 10:14PM
// 输出 Cookie 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToCookieStringWithTimezone(carbon.LosAngeles) // Tuesday, 04-Aug-2020 22:14:15 PDT
// 输出 DayDateTime 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeStringWithTimezone(carbon.LosAngeles) // Tue, Aug 4, 2020 10:14 PM
// 输出 RSS 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRssStringWithTimezone(carbon.LosAngeles) // Tue, 04 Aug 2020 22:14:15 -0700
// 输出 W3C 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToW3cStringWithTimezone(carbon.LosAngeles) // 2020-08-04T22:14:15-07:00

// 输出 RFC822 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc822StringWithTimezone(carbon.LosAngeles) // 04 Aug 20 22:14 PDT
// 输出 RFC822Z 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc822zStringWithTimezone(carbon.LosAngeles) // 04 Aug 20 22:14 -0700
// 输出 RFC850 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc850StringWithTimezone(carbon.LosAngeles) // Tuesday, 04-Aug-20 22:14:15 PDT
// 输出 RFC1036 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc1036StringWithTimezone(carbon.LosAngeles) // Tue, 04 Aug 20 22:14:15 -0700
// 输出 RFC1123 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
carbon.Parse("2020-08-05 13:14:15").ToRfc1123StringWithTimezone() // Tue, 04 Aug 2020 22:14:15 PDT
// 输出 RFC2822 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
carbon.Parse("2020-08-05 13:14:15").ToRfc2822StringWithTimezone(carbon.LosAngeles) // Tue, 04 Aug 2020 22:14:15 -0700
// 输出 RFC3339 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc3339String() // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").ToRfc3339StringWithTimezone(carbon.LosAngeles) // 2020-08-04T22:14:15-07:00
// 输出 RFC7231 格式字符串
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 05:14:15 GMT
carbon.Parse("2020-08-05 13:14:15").ToRfc7231StringWithTimezone(carbon.LosAngeles) // Tue, 04 Aug 2020 22:14:15 GMT

// 输出字符串
fmt.Println(carbon.Parse("2020-08-05")) // 2020-08-05 00:00:00
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
// 输出格式化字符串，Format() 是 ToFormatString() 的简写
carbon.Parse("2020-08-05 13:14:15").ToFormatString("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").ToFormatString("Y年m月d H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM

```
>更多格式化输出符号请查看附录 <a href="#格式化符号表">格式化符号表</a>

##### 时间设置
```go
// 设置时区
carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(carbon.Tokyo).SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15

// 设置语言
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans()) // 1 month ago
carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans()) // 1 月前

// 设置年
carbon.Parse("2019-08-05").SetYear(2020).ToDateString() // 2020-08-05
carbon.Parse("2020-02-29").SetYear(2019).ToDateString() // 2019-03-01

// 设置月
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
carbon.Parse("2020-08-05").SetMonth(2).ToDateString() // 2020-02-05

// 设置日
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 设置时
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 设置分
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 设置秒
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00
```

##### 时间获取
```go
// 获取本年总天数
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 获取本月总天数
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 获取本年第几天
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 获取本年第几周
carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32

// 获取本月第几天(从1开始)
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 获取本月第几周(从1开始)
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 获取本周第几天(从1开始)
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3
carbon.Parse("2020-08-09 13:14:15").Week() // 7

// 获取当前世纪
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// 获取当前年份
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 获取当前季度
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 获取当前月份
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 获取当前周(从0开始)
carbon.Parse("2020-08-05 13:14:15").Week() // 3
carbon.Parse("2020-08-09 13:14:15").Week() // 0
// 获取当前天数
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 获取当前时
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 获取当前分钟
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 获取当前秒钟
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// 获取当前毫秒
carbon.Parse("2020-08-05 13:14:15").Millisecond() // 1596604455000
// 获取当前微秒
carbon.Parse("2020-08-05 13:14:15").Microsecond() // 1596604455000000
// 获取当前纳秒
carbon.Parse("2020-08-05 13:14:15").Nanosecond() // 1596604455000000000

// 获取时区
carbon.SetTimezone(carbon.PRC).Timezone() // PRC
carbon.SetTimezone(carbon.Tokyo).Timezone() // Asia/Tokyo

// 获取当前语言
carbon.Now().Locale() // en
carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

// 获取当前星座(需要先把lang目录复制到项目目录下)
carbon.Now().Constellation() // Leo
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("zh-CN").Constellation() // 狮子座

// 获取年龄
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18

```
> 关于第几周的计算如有疑惑请查看 [ISO8601标准](https://baike.baidu.com/item/ISO%208601/3910715)

##### 农历
```go
// 获取生肖年
carbon.Parse("2020-08-05 13:14:15").AnimalYear() // 鼠
// 获取干支纪年
carbon.Parse("2020-08-05 13:14:15").ToChineseYearString() // 庚子
// 获取干支纪日
carbon.Parse("2020-08-05 13:14:15").ToChineseDayString() // 庚辰

// 是否是鼠年
carbon.Parse("2020-08-05 13:14:15").IsYearOfRat() // true
// 是否是牛年
carbon.Parse("2020-08-05 13:14:15").IsYearOfOx() // false
// 是否是虎年
carbon.Parse("2020-08-05 13:14:15").IsYearOfTiger() // false
// 是否是兔年
carbon.Parse("2020-08-05 13:14:15").IsYearOfRabbit() // false
// 是否是龙年
carbon.Parse("2020-08-05 13:14:15").IsYearOfDragon() // false
// 是否是蛇年
carbon.Parse("2020-08-05 13:14:15").IsYearOfSnake() // false
// 是否是马年
carbon.Parse("2020-08-05 13:14:15").IsYearOfHorse() // false
// 是否是羊年
carbon.Parse("2020-08-05 13:14:15").IsYearOfGoat() // false
// 是否是猴年
carbon.Parse("2020-08-05 13:14:15").IsYearOfMonkey() // false
// 是否是鸡年
carbon.Parse("2020-08-05 13:14:15").IsYearOfRooster() // false
// 是否是狗年
carbon.Parse("2020-08-05 13:14:15").IsYearOfDog() // false
// 是否是猪年
carbon.Parse("2020-08-05 13:14:15").IsYearOfPig() // false
```

##### 星座
```go
// 获取星座
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo

// 是否是白羊座
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// 是否是金牛座
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// 是否是双子座
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// 是否是巨蟹座
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// 是否是狮子座
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// 是否是处女座
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// 是否是天秤座
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// 是否是天蝎座
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// 是否是射手座
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// 是否是摩羯座
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// 是否是水瓶座
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// 是否是双鱼座
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### 数据库支持
> 假设数据表为 users，字段有 id(int)、name(varchar)、age(int)、birthday(datetime)、graduated_at(datetime)、created_at(datetime)、updated_at(datetime)、date_time1(datetime)、date_time2(datetime)、date_time3(datetime)、date_time4(datetime)

###### 定义模型
```go
type UserModel struct {
    ID  int64  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday carbon.ToDateTimeString `json:"birthday"`
    GraduatedAt carbon.ToDateString `json:"graduated_at"`
    CreatedAt carbon.ToTimeString `json:"created_at"`
    UpdatedAt carbon.ToTimestamp `json:"updated_at"`
    DateTime1 carbon.ToTimestampWithSecond `json:"date_time1"`
    DateTime2 carbon.ToTimestampWithMillisecond `json:"date_time2"`
    DateTime3 carbon.ToTimestampWithMicrosecond `json:"date_time3"`
    DateTime4 carbon.ToTimestampWithNanosecond `json:"date_time4"`
}
```

###### 实例化模型
```go
user := UserModel {
    ID: 1153,
    Name: "勾国印",
    Age: 18,
    Birthday: carbon.ToDateTimeString{carbon.Now().SubYears(18)},
    GraduatedAt: carbon.ToDateString{carbon.Parse("2012-09-09")},
    CreatedAt: carbon.ToTimeString{carbon.Now()},
    UpdatedAt: carbon.ToTimestamp{carbon.Now()},
    DateTime1: carbon.ToTimestampWithSecond{carbon.Now()},
    DateTime2: carbon.ToTimestampWithMillisecond{carbon.Now()},
    DateTime3: carbon.ToTimestampWithMicrosecond{carbon.Now()},
    DateTime4: carbon.ToTimestampWithNanosecond{carbon.Now()},
}
```

###### 输出模型字段
```go
user.ID // 1153
user.Name // 勾国印
user.Age // 18
user.Birthday.ToDateTimeString() // 2012-08-05 13:14:15
user.GraduatedAt.ToDateString() // 2012-09-09
user.CreatedAt.ToTimeString() // 13:14:15
user.UpdatedAt.ToTimestamp() // 1596604455
user.DateTime1.ToTimestampWithSecond() // 1596604455
user.DateTime2.ToTimestampWithMillisecond() // 1596604455000
user.DateTime3.ToTimestampWithMicrosecond() // 1596604455000000
user.DateTime4.ToTimestampWithNanosecond() // 1596604455000000000
```

###### JSON 输出模型
```go
data, _ := json.Marshal(&user)
fmt.Print(string(data))
// 输出
{
    "id": 1153,
    "name": "勾国印",
    "age": 18,
    "birthday": "2012-08-05 13:14:15",
    "graduated_at": "2012-09-09",
    "created_at": "13:14:15",
    "updated_at": 1596604455,
    "date_time1": 1596604455,
    "date_time2": 1596604455000,
    "date_time3": 1596604455000000,
    "date_time4": 1596604455000000000,
}
```

###### 输出自定义格式
```go
// 定义输出格式
type ToRssString struct {
    carbon.Carbon
}

// 定义模型
type UserModel struct {
    Birthday carbon.ToRssString `json:"birthday"`
}

// 实例化模型
user := UserModel {
    Birthday: carbon.ToRssString{carbon.Now()},
}

// 重写MarshalJSON方法
func (c ToRssString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToRssString())), nil
}

// json.Marshal(&user)输出
{
    "birthday": "Wed, 05 Aug 2020 13:14:15 +0800",
}
```


##### 国际化
> 需要使用多语言时，请先把lang目录复制到项目目录下

目前支持的语言有
* [简体中文(zh-CN)](./lang/zh-CN.json "简体中文")
* [繁体中文(zh-TW)](./lang/zh-TW.json "繁体中文")
* [英语(en)](./lang/en.json "英语")
* [日语(jp)](./lang/jp.json "日语")

目前支持的方法有
* DiffForHumans()：输出对人类友好的可读格式时间差
* ToMonthString()：输出完整月份字符串
* ToShortMonthString()：输出缩写月份字符串
* ToWeekString()：输出完整星期字符串
* ToShortWeekString()：输出缩写星期字符串
* Constellation()：获取星座

###### 设置区域
```go
// 方式一(推荐)
c := carbon.Now().AddHours(1).SetLocale("zh-CN") 
if c.Error != nil {
    // 错误处理
    log.Fatal(c.Error)
}
c.DiffForHumans() // 1 小时后
c.ToMonthString() // 八月
c.ToShortMonthString() // 8月
c.ToWeekString() // 星期二
c.ToShortWeekString() // 周二
c.Constellation() // 狮子座

// 方式二
lang := NewLanguage()
if err := lang.SetLocale("zh-CN");err != nil {
	// 错误处理
    log.Fatal(err)
}
c.DiffForHumans() // 1 小时后
c.ToMonthString() // 八月
c.ToShortMonthString() // 8月
c.ToWeekString() // 星期二
c.ToShortWeekString() // 周二
c.Constellation() // 狮子座
```
###### 设置目录
```go
lang := NewLanguage()
if err := lang.SetDir("lang");err != nil {
	// 错误处理
    log.Fatal(err)
}
c.DiffForHumans() // 1 hour from now
c.ToMonthString() // August
c.ToShortMonthString() // Aug
c.ToWeekString() // Tuesday
c.ToShortWeekString() // Tue
c.Constellation() // Leo
```

###### 重写部分翻译资源(其余仍然按照指定的 locale 翻译)
```go
lang := NewLanguage()

if err := lang.SetLocale("en");err != nil {
	// 错误处理
    log.Fatal(err)
}

resources := map[string]string {
    "hour":"%dh",
}
lang.SetResources(resources)

carbon.Now().AddYears(1).SetLanguage(lang).DiffForHumans() // 1 year from now
carbon.Now().AddHours(1).SetLanguage(lang).DiffForHumans() // 1h from now
carbon.Now().SetLanguage(lang).ToMonthString() // August
carbon.Now().SetLanguage(lang).ToShortMonthString() // Aug
carbon.Now().SetLanguage(lang).ToWeekString() // Tuesday
carbon.Now().SetLanguage(lang).ToShortWeekString() // Tue
carbon.Now().SetLanguage(lang).Constellation() // Leo
```

###### 重写全部翻译资源(无需指定 locale)
```go
lang := NewLanguage()
resources := map[string]string {
    "months": "January|February|March|April|May|June|July|August|September|October|November|December",
    "months_short": "Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec",
    "weeks": "Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday",
    "weeks_short": "Sun|Mon|Tue|Wed|Thu|Fri|Sat",
    "constellations": "Aries|Taurus|Gemini|Cancer|Leo|Virgo|Libra|Scorpio|Sagittarius|Capricornus|Aquarius|Pisce",
    "year":"1 yr|%d yrs",
    "month":"1 mo|%d mos",
    "week":"%dw",
    "day":"%dd",
    "hour":"%dh",
    "minute":"%dm",
    "second":"%ds",
    "now": "just now",
    "ago":"%s ago",
    "from_now":"in %s",
    "before":"%s before",
    "after":"%s after",
}
lang.SetResources(resources)

carbon.Now().AddYears(1).SetLanguage(lang).DiffForHumans() // in 1 yr
carbon.Now().AddHours(1).SetLanguage(lang).DiffForHumans() // in 1h
carbon.Now().SetLanguage(lang).ToMonthString() // August
carbon.Now().SetLanguage(lang).ToShortMonthString() // Aug
carbon.Now().SetLanguage(lang).ToWeekString() // Tuesday
carbon.Now().SetLanguage(lang).ToShortWeekString() // Tue
carbon.Now().SetLanguage(lang).Constellation() // Leo
```

##### 错误处理
> 如果有多个错误发生，只返回第一个错误信息，前一个错误排除后才返回下一个错误信息

###### 场景一
```go
c := carbon.SetTimezone(PRC).Parse("123456")
if c.Error != nil {
    // 错误处理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 输出
the value "123456" can't parse string as time
```
###### 场景二
```go
c := carbon.SetTimezone("XXXX").Parse("2020-08-05")
if c.Error != nil {
    // 错误处理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 输出
invalid timezone "XXXX", please see the $GOROOT/lib/time/zoneinfo.zip file for all valid timezone
```
###### 场景三
```go
c := carbon.SetTimezone("XXXX").Parse("12345678")
if c.Error != nil {
    // 错误处理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 输出
invalid timezone "XXXX", please see the $GOROOT/lib/time/zoneinfo.zip file for all valid timezone
```
> 建议使用SetTimezone()、Parse()、ParseByFormat()、AddDuration()、SubDuration()、SetLocale()等方法时先进行错误处理判断，除非你能确保传入参数无误
#### 附录
##### <a id="格式化符号表">格式化符号表</a>
| 符号 | 描述 |  长度 | 范围 | 示例 |
| :------------: | :------------: | :------------: | :------------: | :------------: |
| d | 月份中的第几天，有前导零 | 2 | 01-31 | 05 |
| D | 缩写单词表示的周几 | 3 | Mon-Sun | Wed |
| j | 月份中的第几天，没有前导零 | - |1-31 | 5 |
| S | 月份中的第几天，英文缩写后缀，一般和j配合使用 | 2 | st/nd/rd/th | th |
| l | 完整单词表示的周几 | - | Monday-Sunday | Wednesday |
| F | 完整单词表示的月份 | - | January-December | August |
| m | 数字表示的月份，有前导零 | 2 | 01-12 | 08 |
| M | 缩写单词表示的月份 | 3 | Jan-Dec | Aug |
| n | 数字表示的月份，没有前导零 | - | 1-12 | 8 |
| y | 年份，有前导零 | 2 | 00-99 | 20 |
| Y | 年份 | 4 | 0000-9999 | 2020 |
| a | 小写的上下午缩写字母 | 2 | am/pm | pm |
| A | 大写的上下午缩写字母 | 2 | AM/PM | PM |
| g | 小时，12 小时格式，没有前导零 | - | 1-12 | 1 |
| G | 小时，24 小时格式，没有前导零 | - | 0-23 | 15 |
| h | 小时，12 小时格式，有前导零 | 2 | 00-11 | 03 |
| H | 小时，24 小时格式，有前导零 | 2 | 00-23 | 15 |
| i | 分钟，有前导零 | 2 | 01-59 | 14 |
| s | 秒数，有前导零 | 2 | 01-59 | 15 |
| c | ISO8601 格式的日期 | - | - | 2020-08-05T15:19:21+00:00 |
| r | RFC822 格式的日期 | - | - | Thu, 21 Dec 2020 16:01:07 +0200 |
| O | 与格林威治时间相差的小时数 | - | - | +0200 |
| P | 与格林威治时间相差的小时数，小时和分钟之间有冒号分隔 | - | - | +02:00 |
| T | 时区缩写 | - | - | EST |
| W | ISO-8601 格式数字表示的年份中的第几周 | - | 1-52 | 42 |
| N | ISO-8601 格式数字表示的星期中的第几天 | 1 | 1-7 | 6 |
| L | 是否为闰年，如果是闰年为 1，否则为 0 | 1 | 0-1 | 1 |
| U | 秒级时间戳 | 10 | - | 1611818268 |
| u | 毫秒 | - | - | 999 |
| w | 数字表示的周几 | 1 | 0-6 | 6 |
| t | 月份中的总天数 | 2 | 28-31 | 30 |
| z | 年份中的第几天 | - | 0-365 | 15 |
| e | 时区标识 | - | - | America/New_York |

#### 参考项目
* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [uniplaces/carbon](https://github.com/uniplaces/carbon)
* [jinzhu/now](https://github.com/jinzhu/now)
* [araddon/dateparse](https://github.com/araddon/dateparse)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)

#### 相关链接
[更新日志](https://github.com/golang-module/carbon/wiki/%E6%9B%B4%E6%96%B0%E6%97%A5%E5%BF%97)