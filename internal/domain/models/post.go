package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time
	User      User 		`json:"user"`
}

func (p Post) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ID, validation.Required, is.UUID),
		validation.Field(&p.UserID, validation.Required, is.UUID),
		validation.Field(&p.Title, validation.Required, validation.Length(1, 50)),
		validation.Field(&p.Content, validation.Required, validation.Length(1, 10000)),
		validation.Field(&p.CreatedAt, validation.Required),
		validation.Field(&p.UpdatedAt, validation.Required),
	)
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
