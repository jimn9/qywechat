package qyWechatAccountModel

import (
	"encoding/json"
	"workwx/app/models"
	"workwx/pkg/logger"
)

type ScrmQyWechatAccount struct {
	models.BaseModel

	ManagerId               int    `gorm:"column:manager_id;type:int"`
	Corpid                  string `gorm:"column:corpid;type:varchar(255)"`
	Agentid                 string `gorm:"column:agentid;type:varchar(255)"`
	Corpsecret              string `gorm:"column:corpsecret;type:varchar(255)"`
	Name                    string `gorm:"column:name;type:varchar(255)"`
	Logo                    string `gorm:"column:logo;type:varchar(255)"`
	CustomerCorpsecret      string `gorm:"column:customer_corpsecret;type:varchar(255)"`
	ConversationCorpsecret  string `gorm:"column:conversation_corpsecret;type:varchar(255)"`
	ConversationPrivate_pem string `gorm:"column:conversation_private_pem;type:text"`
	CustomerCount           int    `gorm:"column:customer_count;type:int"`
	UserCount               int    `gorm:"column:user_count;type:int"`
	AdminCount              int    `gorm:"column:admin_count;type:int"`
	ServerToken             string `gorm:"column:server_token;type:varchar(255)"`
	ServerEncodingAesKey    string `gorm:"column:server_encoding_aes_key;type:varchar(255)"`
	IsDelete                int    `gorm:"column:is_delete;type:tinyint"`
	DelFollowUser           int    `gorm:"column:del_follow_user;type:tinyint"`
	DelCustomer             int    `gorm:"column:del_customer;type:tinyint"`
	DelCustomerType         int    `gorm:"column:del_customer_type;type:tinyint"`
	DelCustomerAcceptors    string `gorm:"column:del_customer_acceptors;type:json"`
	DelCustomerLimit        int    `gorm:"column:del_customer_limit;type:int"`
}

func (ScrmQyWechatAccount) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "scrm_qy_wechat_account"
}

func (account ScrmQyWechatAccount)GetDelCustomerAcceptors() []string{
	var acceptors []string
	err := json.Unmarshal([]byte(account.DelCustomerAcceptors),&acceptors)
	logger.LogError(err)
	return acceptors
}

