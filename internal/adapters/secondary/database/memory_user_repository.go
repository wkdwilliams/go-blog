package database

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
)

type MemoryUserRepository struct {
	users map[uuid.UUID]*models.User
	mu    sync.RWMutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[uuid.UUID]*models.User),
	}
}

func (ur *MemoryUserRepository) Add(u *models.User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	// If the user ID is not set, generate a new one
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	ur.users[u.ID] = u
	return nil
}

func (ur *MemoryUserRepository) GetById(id uuid.UUID) (*models.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	user, exists := ur.users[id]
	if !exists {
		return nil, ports.ErrRecordNotFound
	}
	return user, nil
}

func (ur *MemoryUserRepository) GetAll() ([]models.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	users := make([]models.User, 0, len(ur.users))
	for _, user := range ur.users {
		users = append(users, *user)
	}
	return users, nil
}

func (ur *MemoryUserRepository) GetByUsername(username string) (*models.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	for _, user := range ur.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (ur *MemoryUserRepository) GetTotalCount() int64 {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	return int64(len(ur.users))
}
