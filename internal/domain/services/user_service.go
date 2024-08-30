package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
)

type IUserService interface {
	CreateAccount(username, password string) (*models.User, error)
	GetById(id uuid.UUID) (*models.User, error)
	GetAll() ([]models.User, error)
}

type UserService struct {
	userRepository ports.UserRepository
}

func (s *UserService) CreateAccount(username, password string) (*models.User, error) {
	user, err := models.NewUser(username, password)

	if err != nil {
		return nil, err
	}

	if err := s.userRepository.Add(user); err != nil {
		return nil, fmt.Errorf("failed to add a user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetById(id uuid.UUID) (*models.User, error) {
	return s.userRepository.GetById(id)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.userRepository.GetAll()
}

func NewUserService(ur ports.UserRepository) *UserService {
	return &UserService{ur}
}
