package cFunc

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"workwx/pkg/logger"
)

func Md5(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func JsonEncode(data interface{}) string {
	res, err := json.Marshal(data)
	logger.LogError(err)
	return string(res)
}

func MustStringKeyMap(v interface{}) map[string]interface{} {
	switch v.(type) {
	case map[string]interface{}:
		res := make(map[string]interface{})
		for k, v := range v.(map[string]interface{}) {
			res[k] = v
		}
		return res
	default:
		return nil
	}
}

func MustStringArray(v interface{}) []string {
	switch v.(type) {
	case []string:
		res := make([]string, 0, 0)
		for _, v := range v.([]string) {
			res = append(res, v)
		}
		return res
	case []interface{}:
		res := make([]string, 0, 0)
		for _, v := range v.([]interface{}) {
			res = append(res, MustString(v))
		}
		return res
	default:
		return nil
	}
}

func MustString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return ""
	}
}

func MustInt(v interface{}) int {
	switch v.(type) {
	case int:
		return v.(int)
	default:
		return 0
	}
}

func MustBool(v interface{}) bool {
	switch v.(type) {
	case bool:
		return v.(bool)
	default:
		return false
	}
}

func MustFloat64(v interface{}) float64 {
	switch v.(type) {
	case float64:
		return v.(float64)
	default:
		return float64(0)
	}
}

func MapMerge(maps ...*map[string]interface{}) {
	mapsLen := len(maps)
	if mapsLen <= 1 {
		return
	}
	firstMap := *maps[0]
	maps = maps[1:mapsLen]
	for _, m := range maps {
		for k, v := range *m {
			firstMap[k] = v
		}
	}

}

//func dump(obj interface{}) error {
//	if obj == nil {
//		fmt.Println("nil")
//		return nil
//	}
//	switch obj.(type) {
//	case bool:
//		fmt.Println(obj.(bool))
//	case int:
//		fmt.Println(obj.(int))
//	case float64:
//		fmt.Println(obj.(float64))
//	case string:
//		fmt.Println(obj.(string))
//	case map[string]interface{}:
//
//	default:
//		return errors.New(
//			fmt.Sprintf("Unsupported type: %v", obj))
//	}
//
//	return nil
//}
