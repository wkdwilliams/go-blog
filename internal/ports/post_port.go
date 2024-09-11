package ports

import (
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

//go:generate mockery --name PostRepository
type PostRepository interface {
	Create(u *models.Post) error
	GetById(id uuid.UUID) (*models.Post, error)
	GetAll() ([]models.Post, error)
	Delete(id uuid.UUID) error
}
