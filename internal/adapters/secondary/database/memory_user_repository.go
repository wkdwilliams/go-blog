package database

import (
	"errors"
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"sync"
)

type MemoryUserRepository struct {
	users map[uuid.UUID]*models.User
	mu    sync.Mutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[uuid.UUID]*models.User),
	}
}

func (ur *MemoryUserRepository) Add(u *models.User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	ur.users[u.ID] = u
	return nil
}

func (ur *MemoryUserRepository) GetById(id uuid.UUID) (*models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	user, exists := ur.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (ur *MemoryUserRepository) GetAll() ([]models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	var users []models.User
	for _, user := range ur.users {
		users = append(users, *user)
	}

	return users, nil
}

func (ur *MemoryUserRepository) GetByUsername(username string) (*models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	for _, user := range ur.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
