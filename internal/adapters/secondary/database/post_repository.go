package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (ur *PostRepository) Create(u *models.Post) error {
	return ur.db.Save(u).Error
}

func (ur *PostRepository) GetById(id uuid.UUID) (*models.Post, error) {
	post := &models.Post{}

	if err := ur.db.Preload("User").First(post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrRecordNotFound
		}
		return nil, err
	}

	return post, nil
}

func (ur *PostRepository) GetAll() ([]models.Post, error) {
	post := []models.Post{}

	if err := ur.db.Order("created_at DESC").Preload("User").Find(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (ur *PostRepository) Delete(id uuid.UUID) error {
	return ur.db.Where("id = ?", id).Delete(&models.Post{}).Error
}
