package qyWechatAccountModel

import (
	"workwx/pkg/general"
	"workwx/pkg/model"
)

// Get 通过 ID 获取文章
func Get(id int) (*ScrmQyWechatAccount, error) {
	var qyWechatAccount *ScrmQyWechatAccount
	if err := model.DB.First(&qyWechatAccount, id).Error; err != nil {
		return qyWechatAccount, err
	}
	return qyWechatAccount, nil
}

func All() ([]*ScrmQyWechatAccount, error) {
	var accounts []*ScrmQyWechatAccount
	if err := model.DB.Where("is_delete", general.FLAG_YES).Find(&accounts).Error; err != nil {
		return accounts, err
	}
	return accounts, nil
}
