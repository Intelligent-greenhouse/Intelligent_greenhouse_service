package model

import "github.com/cypunsource/cypunsource-tool/database"

type UserGreenhouse struct {
	ID           int32 `json:"id" gorm:"column:id; primaryKey"`
	GreenhouseId int32 `json:"greenhouse_id" gorm:"column:greenhouse_id"`
	UserId       int32 `json:"user_id" gorm:"column:user_id"`
	database.DeleteModel
	database.MetadataTimeModel
}

func (UserGreenhouse) TableName() string {
	return "user_greenhouse"
}
