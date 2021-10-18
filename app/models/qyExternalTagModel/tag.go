package qyExternalTagModel

import (
	"time"
	"workwx/app/models"
	group "workwx/app/models/qyExternalTagGroupModel"
)

type QyExternalTag struct {
	models.BaseModel

	QyId       int                      `gorm:"column:qy_id;type:int"`
	GroupId    string                   `gorm:"column:group_id;type:varchar(60)"`
	Group      group.QyExternalTagGroup `gorm:"foreignKey:GroupId;references:GroupId"`
	TagId      string                   `gorm:"column:tag_id;type:varchar(60)"`
	Name       string                   `gorm:"column:name;type:varchar(120)"`
	CreateTime time.Time                `gorm:"column:name;type:datetime"`
	Order      int                      `gorm:"column:order;type:int"`
	IsDelete   int                      `gorm:"column:is_delete;type:tinyint"`
}

func (QyExternalTag) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "scrm_qy_external_tag"
}
