package ports

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
)

type UserRepository interface {
	Add(u *models.User) error
	GetById(id uuid.UUID) (*models.User, error)
	GetAll() ([]models.User, error)
}
