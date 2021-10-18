package models

import (
	"time"
	cFunc "workwx/pkg/commonFunc"
	"workwx/pkg/model"
	"workwx/pkg/types"

	"github.com/golang-module/carbon"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`

	// 支持 gorm 软删除
	// DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" sql:"index"`
}

// GetStringID 获取 ID 的字符串格式
func (a *BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

func (a *BaseModel) CreateAtDate() string {
	return carbon.Time2Carbon(a.CreatedAt).ToDateString()

}

func (a *BaseModel) CreateAtDateTime() string {
	return carbon.Time2Carbon(a.CreatedAt).ToDateTimeString()
}

func (a *BaseModel) UpdateAtDate() string {
	return carbon.Time2Carbon(a.UpdatedAt).ToDateString()

}

func (a *BaseModel) UpdateAtDateTime() string {
	return carbon.Time2Carbon(a.UpdatedAt).ToDateTimeString()
}

func (a *BaseModel) Time2Carbon(t time.Time) carbon.Carbon {
	return carbon.Time2Carbon(t)
}

func (m *BaseModel) UpdatedOrCreate(mod interface{}, wheres, datas map[string]interface{}) int {
	result := map[string]interface{}{}
	var RowsAffected int64
	err := model.DB.Model(mod).Where(wheres).First(&result).Error
	now := time.Now()
	datas["updated_at"] = now
	if err != nil {
		datas["created_at"] = now
		cFunc.MapMerge(&datas, &wheres)
		RowsAffected = model.DB.Model(mod).Create(datas).RowsAffected
	} else {
		//存在
		RowsAffected = model.DB.Model(mod).Where("id = ?", result["ID"]).Updates(datas).RowsAffected
	}
	return int(RowsAffected)
}
