package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID                 uint   `gorm:"primaryKey;autoIncrement"`
	Name               string `gorm:"size:100;not null"`
	ProfileImageUrl    string `gorm:"size:500"`
	PhoneNumber        string `gorm:"index;size:100;not null"`
	Password           string `gorm:"size:500;not null"`
	ForgotPasswordCode uint16
	CreatedAt          int `gorm:"autoCreateTime:mili"`
	UpdatedAt          int `gorm:"autoUpdateTime:mili"`
}

func InsertUserIntoDB(db *gorm.DB, user User) {
	db.Create(&user)
}
