package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"workwx/pkg/types"
)

func Curl(method, url string, params map[string]interface{}) ([]byte, error) {
	method = strings.ToLower(method)
	switch method {
	case "get":
		return get(url, params)
	case "post":
		return post(url, params)
	default:
		return nil, errors.New("unknown method")
	}
}

func get(path string, params map[string]interface{}) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	if params != nil {
		paramsString := mapString2UrlValue(mapInterface2String(params)).Encode()
		path = path + "?" + paramsString
	}
	response, err := client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return parseResponse(response)

}

func post(url string, params map[string]interface{}) ([]byte, error) {
	jsonStr, _ := json.Marshal(params)
	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return parseResponse(response)
}

//统一处理请求结果
func parseResponse(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed , status code is " + types.Int2String(resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("read response failed")
	}
	return body, nil
}

//map[string]interface{}  转为map[string]string
func mapInterface2String(inputData map[string]interface{}) map[string]string {
	outputData := map[string]string{}
	for key, value := range inputData {
		switch value.(type) {
		case string:
			outputData[key] = value.(string)
		case int:
			tmp := value.(int)
			outputData[key] = types.Int2String(tmp)
		case int64:
			tmp := value.(int64)
			outputData[key] = types.Int64ToString(tmp)
		}
	}
	return outputData
}

//map[string]string  转为 url.Values{}
func mapString2UrlValue(inputData map[string]string) url.Values {
	outputData := url.Values{}
	for key, value := range inputData {
		outputData.Set(key, value)
	}
	return outputData
}
