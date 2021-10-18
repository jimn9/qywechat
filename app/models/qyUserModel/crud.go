package qyUserModel

import (
	"workwx/pkg/general"
	"workwx/pkg/logger"
	"workwx/pkg/model"
)

func AccountAllUsers(qy_id int) ([]*User, error) {
	var users []*User
	if err := model.DB.Where("qy_id", qy_id).
		Where("is_delete", general.FLAG_YES).
		Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func UsersUseridToName(qyId int) map[string]string {
	usersGet, err := AccountAllUsers(qyId)
	logger.LogError(err)
	users := make(map[string]string)
	for _, user := range usersGet {
		users[user.Userid] = user.Name
	}
	return users
}
