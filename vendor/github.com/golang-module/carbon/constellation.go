package carbon

import "strings"

// Constellation 获取星座，支持i18n
func (c Carbon) Constellation() string {
	if c.IsZero() {
		return ""
	}
	if len(c.Lang.resources) == 0 && c.Lang.SetLocale(defaultLocale) != nil {
		return ""
	}
	index := -1
	switch {
	case c.Month() == 3 && c.Day() >= 21, c.Month() == 4 && c.Day() <= 19:
		index = 0 // 白羊座
	case c.Month() == 4 && c.Day() >= 20, c.Month() == 5 && c.Day() <= 20:
		index = 1 // 金牛座
	case c.Month() == 5 && c.Day() >= 21, c.Month() == 6 && c.Day() <= 21:
		index = 2 // 双子座
	case c.Month() == 6 && c.Day() >= 22, c.Month() == 7 && c.Day() <= 22:
		index = 3 // 巨蟹座
	case c.Month() == 7 && c.Day() >= 23, c.Month() == 8 && c.Day() <= 22:
		index = 4 // 狮子座
	case c.Month() == 8 && c.Day() >= 23, c.Month() == 9 && c.Day() <= 22:
		index = 5 // 处女座
	case c.Month() == 9 && c.Day() >= 23, c.Month() == 10 && c.Day() <= 23:
		index = 6 // 天秤座
	case c.Month() == 10 && c.Day() >= 24, c.Month() == 11 && c.Day() <= 22:
		index = 7 // 天蝎座
	case c.Month() == 11 && c.Day() >= 23, c.Month() == 12 && c.Day() <= 21:
		index = 8 // 射手座
	case c.Month() == 12 && c.Day() >= 22, c.Month() == 1 && c.Day() <= 19:
		index = 9 // 摩羯座
	case c.Month() == 1 && c.Day() >= 20, c.Month() == 2 && c.Day() <= 18:
		index = 10 // 水瓶座
	case c.Month() == 2 && c.Day() >= 19, c.Month() == 3 && c.Day() <= 20:
		index = 11 // 双鱼座
	default:
		return ""
	}
	if constellations, ok := c.Lang.resources["constellations"]; ok {
		slice := strings.Split(constellations, "|")
		return slice[index]
	}
	return ""
}

// IsAries 是否是白羊座
func (c Carbon) IsAries() bool {
	if c.Month() == 3 && c.Day() >= 21 {
		return true
	}
	if c.Month() == 4 && c.Day() <= 19 {
		return true
	}
	return false
}

// IsTaurus 是否是金牛座
func (c Carbon) IsTaurus() bool {
	if c.Month() == 4 && c.Day() >= 20 {
		return true
	}
	if c.Month() == 5 && c.Day() <= 20 {
		return true
	}
	return false
}

// IsGemini 是否是双子座
func (c Carbon) IsGemini() bool {
	if c.Month() == 5 && c.Day() >= 21 {
		return true
	}
	if c.Month() == 6 && c.Day() <= 21 {
		return true
	}
	return false
}

// IsCancer 是否是巨蟹座
func (c Carbon) IsCancer() bool {
	if c.Month() == 6 && c.Day() >= 22 {
		return true
	}
	if c.Month() == 7 && c.Day() <= 22 {
		return true
	}
	return false
}

// IsLeo 是否是狮子座
func (c Carbon) IsLeo() bool {
	if c.Month() == 7 && c.Day() >= 23 {
		return true
	}
	if c.Month() == 8 && c.Day() <= 22 {
		return true
	}
	return false
}

// IsVirgo 是否是处女座
func (c Carbon) IsVirgo() bool {
	if c.Month() == 8 && c.Day() >= 23 {
		return true
	}
	if c.Month() == 9 && c.Day() <= 22 {
		return true
	}
	return false
}

// IsLibra 是否是天秤座
func (c Carbon) IsLibra() bool {
	if c.Month() == 9 && c.Day() >= 23 {
		return true
	}
	if c.Month() == 10 && c.Day() <= 23 {
		return true
	}
	return false
}

// IsScorpio 是否是天蝎座
func (c Carbon) IsScorpio() bool {
	if c.Month() == 10 && c.Day() >= 24 {
		return true
	}
	if c.Month() == 11 && c.Day() <= 22 {
		return true
	}
	return false
}

// IsSagittarius 是否是射手座
func (c Carbon) IsSagittarius() bool {
	if c.Month() == 11 && c.Day() >= 22 {
		return true
	}
	if c.Month() == 12 && c.Day() <= 21 {
		return true
	}
	return false
}

// IsCapricorn 是否是摩羯座
func (c Carbon) IsCapricorn() bool {
	if c.Month() == 12 && c.Day() >= 22 {
		return true
	}
	if c.Month() == 1 && c.Day() <= 19 {
		return true
	}
	return false
}

// IsAquarius 是否是水瓶座
func (c Carbon) IsAquarius() bool {
	if c.Month() == 1 && c.Day() >= 20 {
		return true
	}
	if c.Month() == 2 && c.Day() <= 18 {
		return true
	}
	return false
}

// IsPisces 是否是双鱼座
func (c Carbon) IsPisces() bool {
	if c.Month() == 2 && c.Day() >= 19 {
		return true
	}
	if c.Month() == 3 && c.Day() <= 20 {
		return true
	}
	return false
}
