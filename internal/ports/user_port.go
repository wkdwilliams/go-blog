package ports

import (
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

type UserRepository interface {
	Create(u *models.User) error
	GetById(id uuid.UUID) (*models.User, error)
	GetAll() ([]models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetTotalCount() int64
}
