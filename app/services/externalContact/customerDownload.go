package externalContact

import (
	"encoding/json"
	"fmt"
	"time"
	"workwx/app/models"
	"workwx/app/models/qyExternalCustomerModel"
	"workwx/app/models/qyExternalCustomerToUserModel"
	"workwx/app/models/qyExternalTagModel"
	"workwx/app/models/qyUserModel"
	"workwx/app/models/qyWechatAccountModel"
	"workwx/app/services/qyWechat/contact"
	cFunc "workwx/pkg/commonFunc"
	"workwx/pkg/general"
	"workwx/pkg/logger"

	"github.com/golang-module/carbon"
)

func CustomerDownloads() {
	accountSync(CustomerDownload)
}

func CustomerDownload(account *qyWechatAccountModel.ScrmQyWechatAccount) {
	if account.ID != 12 {
		return
	}
	qyId := int(account.ID)
	contact := contact.NewExternalContact(account.Corpid, account.Corpsecret)
	resp, err := contact.FollowUser()
	if err != nil {
		return
	}
	if resp.Errcode != 0 {
		//没法获取  跳过
		return
	}

	if len(resp.FollowUser) == 0 {
		//无内部成员
		return
	}

	tags := qyExternalTagModel.CustomerTagFormat(qyId)
	users := qyUserModel.UsersUseridToName(qyId)
	time := time.Now().Unix()
	for _, follow := range resp.FollowUser {
		fmt.Println(users[follow])
		cursor := ""
		for {
			batchDetails, err := contact.BatchDetails(follow, cursor)
			if err != nil {
				//请求不成功
				logger.LogError(err)
				fmt.Println("请求不成功")
				break
			}
			if batchDetails.Errcode != 0 {
				//没法获取  跳过
				fmt.Println("没法获取  跳过")
				break
			}
			externalContactList := batchDetails.ExternalContactList
			fmt.Println(len(externalContactList))
			cursor = batchDetails.NextCursor
			for _, customer := range externalContactList {
				externalContact := customer.ExternalContact
				followInfo := customer.FollowInfo
				customerTagsStr := getTagStr(followInfo.TagID, tags)

				cusWhere := make(map[string]interface{})
				cusDatas := make(map[string]interface{})
				folWhere := make(map[string]interface{})
				folDatas := make(map[string]interface{})

				cusWhere["qy_id"] = account.ID
				folWhere["qy_id"] = account.ID
				cusWhere["external_userid"] = externalContact.ExternalUserid
				folWhere["external_userid"] = externalContact.ExternalUserid
				folWhere["userid"] = follow
				cusDatas["name"] = externalContact.Name
				cusDatas["type"] = externalContact.Type
				cusDatas["gender"] = externalContact.Gender
				cusDatas["unionid"] = externalContact.Unionid
				cusDatas["avatar"] = externalContact.Avatar
				cusDatas["first_join_state"] = followInfo.State
				cusDatas["is_delete"] = general.FLAG_YES
				cusDatas["position"] = externalContact.Position
				cusDatas["corp_name"] = externalContact.CorpName
				cusDatas["corp_full_name"] = externalContact.CorpFullName
				cusDatas["external_profile"] = cFunc.JsonEncode(externalContact.ExternalProfile)
				folDatas["name"] = users[followInfo.Userid]
				folDatas["remark"] = followInfo.Remark
				folDatas["description"] = followInfo.Description
				createtime := int64(followInfo.Createtime)
				if createtime == 0 {
					createtime = time
				}
				folDatas["createtime"] = carbon.CreateFromTimestamp(createtime).ToDateTimeString()
				folDatas["tags"] = customerTagsStr
				folDatas["remark_corp_name"] = followInfo.RemarkCorpName
				folDatas["remark_mobiles"] = cFunc.JsonEncode(followInfo.RemarkMobiles)
				folDatas["add_way"] = followInfo.AddWay
				folDatas["oper_userid"] = followInfo.OperUserid
				folDatas["state"] = followInfo.State
				folDatas["is_delete"] = general.FLAG_NO
				new(models.BaseModel).UpdatedOrCreate(qyExternalCustomerModel.QyExternalCustomer{}, cusWhere, cusDatas)
				new(models.BaseModel).UpdatedOrCreate(qyExternalCustomerToUserModel.QyExternalCustomerToUser{}, folWhere, folDatas)
			}
			if cursor == "" {
				break
			}
		}
	}

}

func getTagStr(tags []string, allTags map[string]string) string {
	res := make([]string, 0)
	for _, tag := range tags {
		if tsgStr, ok := allTags[tag]; ok {
			res = append(res, tsgStr)
		}
	}
	tagsStr, _ := json.Marshal(res)
	return string(tagsStr)
}
