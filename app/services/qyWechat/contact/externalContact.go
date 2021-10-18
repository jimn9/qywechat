package contact

import (
	"encoding/json"
	"errors"
	"workwx/app/services/qyWechat/constant"
	"workwx/pkg/general"
)

const (
	EXTERNAL_CONTACT_BASE_URL = BASE_URL + "/externalcontact"
	GET_CUSTOMER_LIST         = EXTERNAL_CONTACT_BASE_URL + "/list"
	GET_CUSTOMER_DETAIL       = EXTERNAL_CONTACT_BASE_URL + "/get"
	POST_CUSTOMER_REMARK      = EXTERNAL_CONTACT_BASE_URL + "/remark"
	POST_BATCH_USER_DETAIL    = EXTERNAL_CONTACT_BASE_URL + "/batch/get_by_user"
	GET_FOLLOW_USER_LIST      = EXTERNAL_CONTACT_BASE_URL + "/get_follow_user_list"
	POST_GET_USER_BEHAVIOR    = EXTERNAL_CONTACT_BASE_URL + "/get_user_behavior_data"
)

type ExternalContact struct {
	Corp
	conversation_corpsecret string
	customer_corpsecret     string
}

func (e *ExternalContact) GetCustomerCorpsecret() string {
	return e.customer_corpsecret
}

func (e *ExternalContact) SetCustomerCorpsecret(customer_corpsecret string) {
	e.customer_corpsecret = customer_corpsecret
}

func (e *ExternalContact) GetConversationCorpsecret() string {
	return e.conversation_corpsecret
}

func (e *ExternalContact) SetConversationCorpsecret(conversation_corpsecret string) {
	e.conversation_corpsecret = conversation_corpsecret
}

func NewExternalContact(id, secret string) *ExternalContact {
	return &ExternalContact{
		Corp: Corp{
			id:     id,
			secret: secret,
			force:  false,
			types:  "1",
		},
		conversation_corpsecret: "",
		customer_corpsecret:     "",
	}
}

func (e *ExternalContact) FollowUser() (*constant.GetFollowUserList, error) {
	e.SetTypes("2")
	token, err := e.GetToken()
	if err != nil {
		return nil, err
	}
	res := e.Curl(general.REQUEST_GET, e.GeneratePath(GET_FOLLOW_USER_LIST, token), nil)
	if res.Errcode != 0 {
		return nil, errors.New(res.Errmsg)
	}
	list := new(constant.GetFollowUserList)
	err = json.Unmarshal(res.Response, &list)

	return list, err
}

func (e *ExternalContact) BatchDetails(userid, cursor string) (*constant.GetByUser, error) {
	e.SetTypes("2")
	token, err := e.GetToken()
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["userid"] = userid
	params["cursor"] = cursor
	params["limit"] = 100
	res := e.Curl(general.REQUEST_POST, e.GeneratePath(POST_BATCH_USER_DETAIL, token), params)
	if res.Errcode != 0 {
		return nil, errors.New(res.Errmsg)
	}
	result := new(constant.GetByUser)
	err = json.Unmarshal(res.Response, &result)

	return result, err

}
