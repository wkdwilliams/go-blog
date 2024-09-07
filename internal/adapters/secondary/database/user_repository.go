package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Add(u *models.User) error {
	return ur.db.Save(u).Error
}

func (ur *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	user := &models.User{}

	if err := ur.db.First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetAll() ([]models.User, error) {
	user := []models.User{}

	if err := ur.db.Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}

	if err := ur.db.First(user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetTotalCount() int64 {
	var count int64

	ur.db.Find(&models.User{}).Count(&count)

	return count
}
