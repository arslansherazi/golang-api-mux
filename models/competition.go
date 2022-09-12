package models

import (
	"time"

	"gorm.io/gorm"
)

type Competition struct {
	ID           uint64  `gorm:"primaryKey;autoIncrement"`
	User       	 User
	Title        string  `gorm:"size:100;not null" validate:"required,min=10,max=100"`
	Description  string  `gorm:"size:5000;not null" validate:"required,min=100,max=5000"`
	Latitude     float64 `gorm:"not null" validate:"required"`
	Longitude    float64 `gorm:"not null" validate:"required"`
	Address      string  `gorm:"size:1000;not null" validate:"required,max=1000"`
	StartingDate string  `gorm:"not null" validate:"required,min=10,max=10"`
	StartingTime string  `gorm:"not null" validate:"required,min=5,max=5"`
	EndingTime   string
	Images       string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// many-to-many relationship
	User []*User `gorm:"many2many:participant"`
}

func InsertCompetitionIntoDB(db *gorm.DB, competition Competition) error {
	result := db.Create(&competition)
	return result.Error
}

func GetCompetitionImagesData(db *gorm.DB, competitionID uint64) (string, error) {
	var competitionImagesData string
	err := db.Table("competition").Select("images").Where("id= ?", competitionID).Find(&competitionImagesData)
	if err.Error != nil {
		return "", err.Error
	}
	return competitionImagesData, nil
}

func EditCompetition(db *gorm.DB, competition Competition) error {
	db.Model(&competition)
	db.Save(&competition)
	return nil
}

func AddParticipant(db *gorm.DB, participant Participant) error {
	result := db.Create(&participant)
	return result.Error
}

func VerifyParticipant(db *gorm.DB, participant Participant) (bool, error) {
	err := db.Find(&participant)
	if err.Error != nil {
		return false, err.Error
	}
	return true, nil
}
