package mysql

import (
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Add(u models.User) error {
	return nil
}

func (ur *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	user := &models.User{}

	ur.db.First(user, id)

	return user, nil
}

func (ur *UserRepository) GetAll() ([]models.User, error) {
	user := []models.User{}

	if err := ur.db.Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
