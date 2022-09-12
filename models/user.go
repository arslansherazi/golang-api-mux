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

	// many-to-many relationship
	Competition []*Competition `gorm:"many2many:participant"`
}

func InsertUserIntoDB(db *gorm.DB, user User) {
	db.Create(&user)
}

func GetUserData(db *gorm.DB, phoneNumber string) User {
	var user User
	db.Where("phone_number = ?", phoneNumber).First(&user)
	return user
}

func ValidatePhoneNumber(db *gorm.DB, phoneNumber string) (bool, error) {
	var id int
	err := db.Table("users").Select("id").Where("phone_number = ?", phoneNumber).Find(&id)
	if err.Error != nil {
		return false, err.Error
	} else if id == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
