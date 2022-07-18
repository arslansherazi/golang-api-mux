package models

import (
	"time"

	"gorm.io/datatypes"
)

type Competiton struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	UserID       int            `gorm:"not null" validate:"required"`
	Title        string         `gorm:"size:100;not null" validate:"required,min=10,max=100"`
	Description  string         `gorm:"size:5000;not null" validate:"required,min=100,max=5000"`
	Latitude     float64        `gorm:"not null" validate:"required"`
	Longitude    float64        `gorm:"not null" validate:"required"`
	Address      string         `gorm:"size:1000;not null" validate:"required,max=1000"`
	StartingDate datatypes.Date `gorm:"not null" validate:"required"`
	EndingDate   datatypes.Date
	StartingTime time.Time `gorm:"not null" validate:"required"`
	EndingTime   time.Time
	Images       datatypes.JSON
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
