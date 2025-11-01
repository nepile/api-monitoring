package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrimaryKey struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
}

func (u *PrimaryKey) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
