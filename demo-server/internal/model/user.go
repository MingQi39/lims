package model

import "time"

type User struct {
	ID           uint64    `gorm:"primaryKey"`
	Username     string    `gorm:"size:64;uniqueIndex;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	DisplayName  string    `gorm:"size:128"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
