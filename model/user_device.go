package model

import (
	"github.com/cypunsource/cypunsource-tool/database"
)

type UserDevice struct {
	ID       int32 `json:"id" gorm:"colum:id; primaryKey"`
	UserId   int32 `json:"userId" gorm:"colum:user_id"`
	DeviceId int32 `json:"deviceId" gorm:"colum:device_id"`
	database.DeleteModel
	database.MetadataTimeModel
}

func (UserDevice) TableName() string {
	return "user_device"
}
