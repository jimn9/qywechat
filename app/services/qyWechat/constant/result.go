package constant

type Result struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type GetFollowUserList struct {
	Result
	FollowUser []string `json:"follow_user"`
}

type GetToken struct {
	Result
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type externalContact struct {
	ExternalUserid  string `json:"external_userid"`
	Name            string `json:"name"`
	Position        string `json:"position"`
	Avatar          string `json:"avatar"`
	CorpName        string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	Type            int    `json:"type"`
	Gender          int    `json:"gender"`
	Unionid         string `json:"unionid"`
	ExternalProfile struct {
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}

type followInfo struct {
	Userid         string   `json:"userid"`
	Remark         string   `json:"remark"`
	Description    string   `json:"description"`
	Createtime     int      `json:"createtime"`
	RemarkCorpName string   `json:"remark_corp_name"`
	RemarkMobiles  []string `json:"remark_mobiles"`
	OperUserid     string   `json:"oper_userid"`
	AddWay         int      `json:"add_way"`
	State          string   `json:"state"`
}

type followInfoTagTypeOne struct {
	followInfo
	TagID []string `json:"tag_id"`
}

type followInfoTagTypeTwo struct {
	followInfo
	Tags []struct {
		GroupName string `json:"group_name"`
		TagName   string `json:"tag_name"`
		TagID     string `json:"tag_id"`
		Type      int    `json:"type"`
	} `json:"tags"`
}

type GetByUser struct {
	Result
	ExternalContactList []struct {
		ExternalContact externalContact      `json:"external_contact"`
		FollowInfo      followInfoTagTypeOne `json:"follow_info"`
	} `json:"external_contact_list"`
	NextCursor string `json:"next_cursor"`
}
