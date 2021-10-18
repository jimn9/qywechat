package carbon

import (
	"math"
)

var (
	// 生肖
	symbolicAnimals = [12]string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}

	// 干支
	stemsAndBranches = [60]string{
		"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉", // 1-10
		"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未", // 11-20
		"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳", // 21-30
		"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯", // 31-40
		"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑", // 41-50
		"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥", // 51-60
	}
)

// AnimalYear 获取生肖年
func (c Carbon) AnimalYear() string {
	if c.IsZero() {
		return ""
	}
	return symbolicAnimals[c.Year()%MonthsPerYear]
}

// Todo ToLunarYear 获取农历年
func (c Carbon) LunarYear() int {
	return 2018
}

// Todo ToLunarMonth 获取农历月
func (c Carbon) LunarMonth() int {
	return 11
}

// Todo LunarDay 获取农历日
func (c Carbon) LunarDay() int {
	return 25
}

// Todo ToLunarYearString 输出农历年
func (c Carbon) ToLunarYearString() string {
	return "二零一八"
}

// Todo ToLunarDayString 输出农历月
func (c Carbon) ToLunarMonthString() string {
	return "十一"
}

// Todo ToLunarDayString 输出农历日
func (c Carbon) ToLunarDayString() string {
	return "廿五"
}

// Todo ToLunarDateString 输出农历日期
func (c Carbon) ToLunarDateString() string {
	return "二零一八年十一月廿五"
}

// Todo ToShortLunarDateString 输出简写农历日期
func (c Carbon) ToShortLunarDateString() string {
	return "20181125"
}

// ToChineseYearString 输出天干地支纪年
func (c Carbon) ToChineseYearString() string {
	if c.IsZero() {
		return ""
	}
	index := c.Year() % 60
	if index < 60 {
		index -= 3
	}
	if index <= 0 {
		index += 60
	}
	return stemsAndBranches[index-1]
}

// Todo ToChineseMonthString 输出天干地支纪月
func (c Carbon) ToChineseMonthString() string {
	if c.IsZero() {
		return ""
	}
	return ""
}

// ToChineseDayString 输出天干地支纪日(根据高氏日柱公式v3版)
func (c Carbon) ToChineseDayString() string {
	if c.IsZero() {
		return ""
	}
	s := c.Year() % 100
	u, d := s%4, c.Day()
	// 计算世纪常数x
	x := (44*(c.Century()-1) + (c.Century()-1)/4 + 9) % 60
	// 计算月基数m
	m := (int(math.Pow(-1, float64(c.Month())))+1)/2*30 + (3*c.Month()-7)/5
	index := (s/4*6 + 5*(s/4*3+u) + m + d + x) % 60
	return stemsAndBranches[index-1]
}

// Todo ToChineseHourString 输出天干地支日期
func (c Carbon) ToChineseDateString() string {
	if c.IsZero() {
		return ""
	}
	return ""
}

// IsYearOfRat 是否是鼠年
func (c Carbon) IsYearOfRat() bool {
	if c.Year()%MonthsPerYear == 4 {
		return true
	}
	return false
}

// IsYearOfOx 是否是牛年
func (c Carbon) IsYearOfOx() bool {
	if c.Year()%MonthsPerYear == 5 {
		return true
	}
	return false
}

// IsYearOfTiger 是否是虎年
func (c Carbon) IsYearOfTiger() bool {
	if c.Year()%MonthsPerYear == 6 {
		return true
	}
	return false
}

// IsYearOfRabbit 是否是兔年
func (c Carbon) IsYearOfRabbit() bool {
	if c.Year()%MonthsPerYear == 7 {
		return true
	}
	return false
}

// IsYearOfDragon 是否是龙年
func (c Carbon) IsYearOfDragon() bool {
	if c.Year()%MonthsPerYear == 8 {
		return true
	}
	return false
}

// IsYearOfSnake 是否是蛇年
func (c Carbon) IsYearOfSnake() bool {
	if c.Year()%MonthsPerYear == 9 {
		return true
	}
	return false
}

// IsYearOfHorse 是否是马年
func (c Carbon) IsYearOfHorse() bool {
	if c.Year()%MonthsPerYear == 10 {
		return true
	}
	return false
}

// IsYearOfGoat 是否是羊年
func (c Carbon) IsYearOfGoat() bool {
	if c.Year()%MonthsPerYear == 11 {
		return true
	}
	return false
}

// IsYearOfMonkey 是否是猴年
func (c Carbon) IsYearOfMonkey() bool {
	if c.Year()%MonthsPerYear == 0 {
		return true
	}
	return false
}

// IsYearOfRooster 是否是鸡年
func (c Carbon) IsYearOfRooster() bool {
	if c.Year()%MonthsPerYear == 1 {
		return true
	}
	return false
}

// IsYearOfDog 是否是狗年
func (c Carbon) IsYearOfDog() bool {
	if c.Year()%MonthsPerYear == 2 {
		return true
	}
	return false
}

// IsYearOfPig 是否是猪年
func (c Carbon) IsYearOfPig() bool {
	if c.Year()%MonthsPerYear == 3 {
		return true
	}
	return false
}
