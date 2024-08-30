package ports

import (
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

type PostRepository interface {
	Add(u models.Post) error
	GetById(id uuid.UUID) (*models.Post, error)
	GetAll() ([]models.Post, error)
}
