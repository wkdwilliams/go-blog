package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:VARCHAR(36);primary_key;"`
	Username  string    `gorm:"index"`
	CreatedAt time.Time
	UpdateAt  time.Time
	Posts     []Post `gorm:"foreignKey:UserID"`
}

func NewUser(username string) User {
	return User{
		ID:       uuid.New(),
		Username: username,
	}
}
