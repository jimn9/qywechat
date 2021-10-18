package qyExternalTagModel

import (
	cFunc "workwx/pkg/commonFunc"
	"workwx/pkg/general"
	"workwx/pkg/logger"
	"workwx/pkg/model"
)

func AllWithGroup(qy_id int) ([]*QyExternalTag, error) {
	var tags []*QyExternalTag
	if err := model.DB.Preload("Group").
		Where("qy_id", qy_id).
		Where("is_delete", general.FLAG_YES).
		Find(&tags).Error;
		err != nil {
		return tags, err
	}
	return tags, nil
}

func CustomerTagFormat(qyId int) map[string]string {
	tagsGet, err := AllWithGroup(qyId)
	logger.LogError(err)
	tags := make(map[string]string)
	for _, tag := range tagsGet {
		tagMap := make(map[string]string)
		tagMap["type"] = "1"
		tagMap["group_name"] = tag.Group.GroupName
		tagMap["tag_id"] = tag.TagId
		tagMap["name"] = tag.Name
		tags[tag.TagId] = cFunc.JsonEncode(tagMap)
	}
	return tags
}
