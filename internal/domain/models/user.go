package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/pkg/hashing"
)

type User struct {
	ID        uuid.UUID `gorm:"type:VARCHAR(16);primary_key;"`
	Username  string    `gorm:"index,unique"`
	Password  string
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
	Posts     []Post `gorm:"foreignKey:UserID"`
}

func NewUser(username, password, name string) (*User, error) {
	hash, err := hashing.HashPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:        uuid.New(),
		Username:  username,
		Password:  hash,
		Name:      name,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}, nil
}
