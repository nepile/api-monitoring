package models

type User struct {
	PrimaryKey
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Timestamp
}
