package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
	"github.com/wkdwilliams/go-blog/pkg/hashing"
)

type IUserService interface {
	CreateAccount(username, password string) (*models.User, error)
	GetById(id uuid.UUID) (*models.User, error)
	GetAll() ([]models.User, error)
	GetByUsername(username string) (*models.User, error)
	Authenticate(username, password string) (*models.User, error)
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

func (s *UserService) GetByUsername(username string) (*models.User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s *UserService) GetById(id uuid.UUID) (*models.User, error) {
	return s.userRepository.GetById(id)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.userRepository.GetAll()
}

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.userRepository.GetByUsername(username)

	if err != nil {
		return nil, err
	}

	if correct := hashing.VerifyPassword(password, user.Password); !correct {
		return nil, errors.New("incorrect username or password")
	}

	return user, nil
}

func NewUserService(ur ports.UserRepository) *UserService {
	return &UserService{ur}
}
