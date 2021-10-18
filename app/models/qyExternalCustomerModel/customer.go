package qyExternalCustomerModel

import (
	"workwx/app/models"
)

type QyExternalCustomer struct {
	models.BaseModel

	QyId            int    `gorm:"column:qy_id;type:int"`
	ExternalUserid  string `gorm:"column:external_userid;type:varchar(50)"`
	Name            string `gorm:"column:name;type:varchar(60)"`
	Avatar          string `gorm:"column:avatar;type:varchar(255)"`
	Type            int    `gorm:"column:type;type:tinyint"`
	Gender          int    `gorm:"column:gender;type:tinyint"`
	Unionid         string `gorm:"column:unionid;type:varchar(255)"`
	Position        string `gorm:"column:position;type:varchar(120)"`
	CorpName        string `gorm:"column:corp_name;type:varchar(120)"`
	CorpFullName    string `gorm:"column:corp_full_name;type:varchar(120)"`
	ExternalProfile string `gorm:"column:external_profile;type:json"`
	FirstJoinState  string `gorm:"column:first_join_state;type:varchar(255)"`
	IsDelete        int    `gorm:"column:is_delete;type:tinyint"`
	IsPermit        int    `gorm:"column:is_permit;type:tinyint"`
}

func (QyExternalCustomer) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "scrm_qy_external_customer"
}
