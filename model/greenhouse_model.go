package model

import "github.com/cypunsource/cypunsource-tool/database"

type Greenhouse struct {
	ID       int32  `json:"id" gorm:"column:id; primaryKey"`
	Username string `json:"name" gorm:"column:name"`
	Pos      string `json:"pos" gorm:"column:pos"`
	Size     int32  `json:"size" gorm:"colum:size"`
	Des      string `json:"des" gorm:"colum:des"`
	database.DeleteModel
	database.MetadataTimeModel
}

func (Greenhouse) TableName() string {
	return "greenhouse"
}
