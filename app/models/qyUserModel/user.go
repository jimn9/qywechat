package qyUserModel

import (
	"workwx/app/models"
)

type User struct {
	models.BaseModel

	QyId             int    `gorm:"column:qy_id;type:int"`
	Userid           string `gorm:"column:userid;type:varchar(50)"`
	Name             string `gorm:"column:name;type:varchar(60)"`
	EnglishName      string `gorm:"column:english_name;type:varchar(255)"`
	HideMobile       int    `gorm:"column:hide_mobile;type:tinyint"`
	Mobile           string `gorm:"column:mobile;type:varchar(20)"`
	Position         string `gorm:"column:mobile;type:varchar(60)"`
	Gender           int    `gorm:"column:gender;type:tinyint"`
	Email            string `gorm:"column:email;type:varchar(60)"`
	Avatar           string `gorm:"column:avatar;type:varchar(255)"`
	ThumbAvatar      string `gorm:"column:thumb_avatar;type:varchar(255)"`
	Telephone        string `gorm:"column:telephone;type:varchar(30)"`
	Alias            string `gorm:"column:alias;type:varchar(60)"`
	Extattr          string `gorm:"column:extattr;type:json"`
	Status           int    `gorm:"column:status;type:tinyint"`
	QrCode           string `gorm:"column:qr_code;type:varchar(255)"`
	Address          string `gorm:"column:address;type:varchar(255)"`
	OpenUserid       string `gorm:"column:open_userid;type:varchar(255)"`
	ExternalPosition string `gorm:"column:external_position;type:varchar(255)"`
	ExternalProfile  string `gorm:"column:external_profile;type:json"`
	IsPermit         int    `gorm:"column:is_permit;type:tinyint"`
	Enable           int    `gorm:"column:enable;type:int"`
	IsLeader         int    `gorm:"column:isleader;type:int"`
	MainDepartment   int    `gorm:"column:main_department;type:int"`
	IsDelete         int    `gorm:"column:is_delete;type:tinyint"`
}

func (User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "scrm_qy_user"
}
