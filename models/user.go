package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID                 uint   `gorm:"primaryKey;autoIncrement"`
	Name               string `gorm:"size:100;not null" validate:"required,min=3,max=100"`
	ProfileImageUrl    string `gorm:"size:500"`
	PhoneNumber        string `gorm:"index;size:100;not null" validate:"required,min=4,max=20"`
	Password           string `gorm:"size:500;not null" validate:"required,min=8,max=16"`
	ForgotPasswordCode uint16
	CreatedAt          int `gorm:"autoCreateTime:mili"`
	UpdatedAt          int `gorm:"autoUpdateTime:mili"`
}

func InsertUserIntoDB(db *gorm.DB, user User) {
	db.Create(&user)
}
