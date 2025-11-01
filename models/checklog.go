package models

import (
	"time"

	"github.com/google/uuid"
)

type CheckLog struct {
	PrimaryKey
	EndpointID   uuid.UUID `gorm:"type:uuid;index"`
	StatusCode   int
	ResponseTime float64
	CheckedAt    time.Time
}
