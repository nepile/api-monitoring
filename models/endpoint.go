package models

import "github.com/google/uuid"

type Endpoint struct {
	PrimaryKey
	UserID         uuid.UUID `gorm:"type:uuid;index"`
	URL            string    `gorm:"not null"`
	ExpectedStatus int       `gorm:"default:200"`
	CheckInterval  int       `gorm:"default:60"`
	Timestamp
}
