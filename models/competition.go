package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Competition struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	UserID       int     `gorm:"not null" validate:"required"`
	Title        string  `gorm:"size:100;not null" validate:"required,min=10,max=100"`
	Description  string  `gorm:"size:5000;not null" validate:"required,min=100,max=5000"`
	Latitude     float64 `gorm:"not null" validate:"required"`
	Longitude    float64 `gorm:"not null" validate:"required"`
	Address      string  `gorm:"size:1000;not null" validate:"required,max=1000"`
	StartingDate string  `gorm:"not null" validate:"required,min=10,max=10"`
	StartingTime string  `gorm:"not null" validate:"required,min=5,max=5"`
	EndingTime   string
	Images       datatypes.JSON
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func InsertCompetitionIntoDB(db *gorm.DB, competition Competition) error {
	result := db.Create(&competition)
	return result.Error
}
