package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts     []Post
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required, is.UUID),
		validation.Field(&u.Username, validation.Required, validation.Length(1, 50)),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Name, validation.Required, validation.Length(1, 30)),
		validation.Field(&u.CreatedAt, validation.Required),
		validation.Field(&u.UpdatedAt, validation.Required),
	)
}

func NewUser(username, password, name string) *User {
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
