package model

import "github.com/cypunsource/cypunsource-tool/database"

type User struct {
	ID       int32  `json:"id" gorm:"column:id; primaryKey"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	IsAdmin  bool   `json:"isAdmin" gorm:"colum:is_admin"`
	database.DeleteModel
	database.MetadataTimeModel
}

func (User) TableName() string {
	return "user"
}
