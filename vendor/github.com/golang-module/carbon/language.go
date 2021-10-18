package carbon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	// 默认目录
	defaultDir = "lang"
	// 默认区域
	defaultLocale = "en"
)

type Language struct {
	dir       string            // 目录
	locale    string            // 区域
	resources map[string]string // 资源
}

// NewLanguage 初始化 Language 实例
func NewLanguage() *Language {
	return &Language{
		dir:       defaultDir,
		locale:    defaultLocale,
		resources: make(map[string]string),
	}
}

// SetLocale 设置区域
func (lang *Language) SetLocale(locale string) error {
	if len(lang.resources) != 0 {
		return nil
	}
	fileName := lang.dir + string(os.PathSeparator) + locale + ".json"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.New("invalid locale \"" + locale + "\", please see the " + lang.dir + " directory for all valid locale")
	}
	if err := json.Unmarshal(bytes, &lang.resources); err != nil {
		return err
	}
	lang.locale = locale
	return nil
}

// SetDir 设置目录
func (lang *Language) SetDir(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		return errors.New("invalid directory \"" + dir + "\", please make sure the directory exists")
	}
	lang.dir = dir
	return nil
}

// SetResources 设置资源
func (lang *Language) SetResources(resources map[string]string) {
	if len(lang.resources) == 0 {
		lang.resources = resources
		return
	}
	for k, v := range resources {
		if _, ok := lang.resources[k]; ok {
			lang.resources[k] = v
		}
	}
}

// translate 翻译转换
func (lang *Language) translate(unit string, diff int64) string {
	if len(lang.resources) == 0 && lang.SetLocale(defaultLocale) != nil {
		return ""
	}
	slice := strings.Split(lang.resources[unit], "|")
	if len(slice) == 1 {
		return strings.Replace(slice[0], "%d", strconv.FormatInt(diff, 10), 1)
	}
	if diff > 1 {
		return strings.Replace(slice[1], "%d", strconv.FormatInt(diff, 10), 1)
	}
	return slice[0]
}
