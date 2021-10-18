package contact

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
	"workwx/app/services/qyWechat/constant"
	cFunc "workwx/pkg/commonFunc"
	"workwx/pkg/curl"
	"workwx/pkg/logger"
	"workwx/pkg/redis"
	"workwx/pkg/types"
)

const (
	BASE_URL         = "https://qyapi.weixin.qq.com/cgi-bin"
	ACCESS_TOKEN_URL = BASE_URL + "/gettoken"
)

type Corp struct {
	id     string
	secret string
	force  bool   //强制获取token  默认为false
	types  string //token类型  默认为1
}

type wechatResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Response []byte
}

type redisToken struct {
	Status int    `json:"status"`
	Errmsg string `json:"errmsg"`
	Token  string `json:"access_token"`
}

func (w *wechatResponse) DefaultErr() *wechatResponse {
	w.Errcode = 99999
	w.Errmsg = "request wechat failed"
	return w
}

func (w *wechatResponse) DecodeErr() *wechatResponse {
	w.Errcode = 99998
	w.Errmsg = "Result Decode failed"
	return w
}

func (w *wechatResponse) OtherErr(code int, msg string) *wechatResponse {
	w.Errcode = code
	w.Errmsg = msg
	return w
}

func NewWechatResponse() *wechatResponse {
	return new(wechatResponse)
}

func NewCorp(id, secret string) *Corp {
	return &Corp{
		id:     id,
		secret: secret,
		force:  false,
		types:  "1",
	}
}

func (c *Corp) GetId() string {
	return c.secret
}

func (c *Corp) GetSecret() string {
	return c.secret
}

func (c *Corp) GetTypes() string {
	return c.types
}

func (c *Corp) GetForce() bool {
	return c.force
}

func (c *Corp) SetSecret(secret string) {
	c.secret = secret
}

func (c *Corp) SetId(id string) {
	c.id = id
}

func (c *Corp) SetTypes(types string) {
	c.types = types
}

func (c *Corp) SetForce(force bool) {
	c.force = force
}

func (c *Corp) Curl(method, path string, params map[string]interface{}) *wechatResponse {
	wechatResponse := NewWechatResponse()
	response, err := curl.Curl(method, path, params)
	if err != nil {
		logger.LogError(err)
		return wechatResponse.DefaultErr()
	}

	err = json.Unmarshal(response, &wechatResponse)
	if err != nil {
		logger.LogError(err)
		return wechatResponse.DecodeErr()
	}

	if wechatResponse.Errcode == 40014 { //这个错误码不会在生成token的时候返回
		//这个时候需要重试  path里面会有access_token   或者params里面   这个token需要替换掉
		//重获token
		token, err := c.GetToken()
		if err != nil {
			return wechatResponse.OtherErr(99997, err.Error())
		}

		//替换token
		i := strings.Index(path, "?access_token=")
		if i == -1 {
			params["access_token"] = token
		} else {
			path = path[0:i] + "?access_token=" + token
		}

		//重新获取结果
		response, err = curl.Curl(method, path, params)
		if err != nil {
			logger.LogError(err)
			return wechatResponse.DefaultErr()
		}
		wechatResponse = NewWechatResponse()
		err = json.Unmarshal(response, &wechatResponse)
		if err != nil {
			logger.LogError(err)
			return wechatResponse.DecodeErr()
		}
	}
	wechatResponse.Response = response
	return wechatResponse
}

func (c *Corp) GetToken() (string, error) {
	rKey := "Token:work_access_token:" + c.id + "_" + cFunc.Md5(c.types+"_"+c.secret)
	res := redis.DB.Get(rKey)
	if res.Err() == nil && c.force == false {
		token := redisToken{}
		json.Unmarshal([]byte(res.Val()), &token)
		return token.Token, nil
	}

	params := make(map[string]interface{})
	params["corpid"] = c.id
	params["corpsecret"] = c.secret
	response := c.Curl("get", ACCESS_TOKEN_URL, params)
	if response.Errcode != 0 {
		err := errors.New(types.Int2String(response.Errcode) + "--" + response.Errmsg)
		logger.LogError(err)
		return "", err
	}
	resToken := constant.GetToken{}
	err := json.Unmarshal(response.Response, &resToken)
	if err != nil {
		logger.LogError(err)
		return "", err
	}

	result := make(map[string]interface{})
	result["status"] = resToken.Errcode
	result["errmsg"] = resToken.Errmsg
	result["access_token"] = resToken.AccessToken
	resultMarshal, _ := json.Marshal(result)
	redis.DB.SetNX(rKey, resultMarshal, 2*time.Hour)
	return resToken.AccessToken, nil
}

func (c *Corp) GeneratePath(path, token string) string {
	return path + "?access_token=" + token
}
