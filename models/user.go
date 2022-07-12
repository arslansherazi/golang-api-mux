package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                 uint   `gorm:"primaryKey;autoIncrement"`
	Name               string `gorm:"size:100;not null" validate:"required,min=3,max=100"`
	ProfileImageUrl    string `gorm:"size:500"`
	PhoneNumber        string `gorm:"index;size:100;not null" validate:"required,min=4,max=20"`
	Password           string `gorm:"size:500;not null" validate:"required,min=8,max=16"`
	ForgotPasswordCode uint16
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func InsertUserIntoDB(db *gorm.DB, user User) {
	db.Create(&user)
}
