package qyExternalCustomerToUserModel

import (
	"time"
	"workwx/app/models"
)

type QyExternalCustomerToUser struct {
	models.BaseModel

	QyId           int       `gorm:"column:qy_id;type:int"`
	ExternalUserid string    `gorm:"column:external_userid;type:varchar(50)"`
	Userid         string    `gorm:"column:userid;type:varchar(50)"`
	Name           string    `gorm:"column:name;type:varchar(255)"`
	Remark         string    `gorm:"column:remark;type:varchar(120)"`
	Description    string    `gorm:"column:description;type:varchar(255)"`
	Createtime     time.Time `gorm:"column:createtime;type:datetime"`
	DelTime        time.Time `gorm:"column:del_time;type:datetime"`
	CusDelUser     int       `gorm:"column:cus_del_user;type:tinyint"`
	UserDelTime    time.Time `gorm:"column:user_del_time;type:datetime"`
	UserDelCus     int       `gorm:"column:user_del_cus;type:tinyint"`
	Tags           string    `gorm:"column:tags;type:json"`
	RemarkCorpName string    `gorm:"column:remark_corp_name;type:varchar(120)"`
	RemarkMobiles  string    `gorm:"column:remark_mobiles;type:varchar(255)"`
	AddWay         int       `gorm:"column:add_way;type:int"`
	OperUserid     string    `gorm:"column:oper_userid;type:varchar(50)"`
	State          string    `gorm:"column:state;type:varchar(100)"`
	IsDelete       int       `gorm:"column:is_delete;type:tinyint"`
}

func (QyExternalCustomerToUser) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "scrm_qy_customer_to_user_tags"
}
