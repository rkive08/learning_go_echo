package models

import "time"

type PasswordReset struct {
	Email     string `gorm:"index"`
	Token     string `gorm:"size:255"`
	CreatedAt time.Time
}
