package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"unique;size:36;not null"`
	Admin     bool      `gorm:"not null"`
	TokenHash string    `gorm:"size:60;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
