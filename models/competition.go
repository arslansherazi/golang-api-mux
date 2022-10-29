package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Competition struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement"`
	UserID       uint64
	Title        string  `gorm:"size:100;not null" validate:"required,min=10,max=100"`
	Description  string  `gorm:"size:5000;not null" validate:"required,min=100,max=5000"`
	Latitude     float64 `gorm:"not null" validate:"required"`
	Longitude    float64 `gorm:"not null" validate:"required"`
	Address      string  `gorm:"size:1000;not null" validate:"required,max=1000"`
	StartingDate string  `gorm:"not null" validate:"required,min=10,max=10"`
	StartingTime string  `gorm:"not null" validate:"required,min=5,max=5"`
	EndingTime   string
	Images       pq.StringArray `gorm:"type:text[]"`
	CreatedAt    uint64         `gorm:"autoCreateTime"`
	UpdatedAt    uint64         `gorm:"autoUpdateTime:milli"`
}

func InsertCompetitionIntoDB(db *gorm.DB, competition Competition) error {
	result := db.Create(&competition)
	return result.Error
}

func GetCompetitionImagesURLs(db *gorm.DB, competitionID uint64) ([]string, error) {
	var competitionImagesURLs []string
	err := db.Table("competition").Select("images").Where("id= ?", competitionID).Find(&competitionImagesURLs)
	if err.Error != nil {
		return nil, err.Error
	}
	return competitionImagesURLs, nil
}

func EditCompetition(db *gorm.DB, competition Competition) error {
	db.Model(&competition)
	db.Save(&competition)
	return nil
}

func AddParticipant(db *gorm.DB, userID uint64, competitionID uint64) error {
	var participant User
	var competition []Competition
	err := db.Table("competition").Where("id= ?", competitionID).Find(&competition)
	if err != nil {
		return err.Error
	}
	err = db.Table("user").Where("id= ?", userID).Find(&participant)
	if err != nil {
		return err.Error
	}
	participant.ParticipationCompetitions = competition
	result := db.Create(&participant)
	return result.Error
}
