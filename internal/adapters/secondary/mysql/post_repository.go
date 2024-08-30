package mysql

import (
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (ur *PostRepository) Add(u models.Post) error {
	return ur.db.Save(u).Error
}

func (ur *PostRepository) GetById(id uuid.UUID) (*models.Post, error) {
	post := &models.Post{}

	if err := ur.db.Preload("User").First(post, id).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (ur *PostRepository) GetAll() ([]models.Post, error) {
	post := []models.Post{}

	if err := ur.db.Preload("User").Find(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}
