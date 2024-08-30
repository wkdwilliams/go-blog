package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:VARCHAR(36);primary_key;"`
	UserID    uuid.UUID `gorm:"type:VARCHAR(36);"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	// User      User
}

func NewPost(title, content string, userId uuid.UUID) Post {
	return Post{
		ID:        uuid.New(),
		UserID:    userId,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
