package models

import (
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement"`
	UserID        int64  `gorm:"not null" validate:"required"`
	CompetitionID int64  `gorm:"not null" validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
