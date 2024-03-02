package model

import "github.com/cypunsource/cypunsource-tool/database"

type GreenhouseDevice struct {
	ID           int32 `json:"id" gorm:"column:id; primaryKey"`
	GreenhouseId int32 `json:"greenhouse_id" gorm:"column:greenhouse_id"`
	DeviceId     int32 `json:"device_id" gorm:"column:device_id"`
	database.DeleteModel
	database.MetadataTimeModel
}

func (GreenhouseDevice) TableName() string {
	return "greenhouse_device"
}
