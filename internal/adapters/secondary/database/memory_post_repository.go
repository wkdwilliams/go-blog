package mysql

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

type MemoryPostRepository struct {
	mu    sync.RWMutex
	posts map[uuid.UUID]models.Post
}

func NewMemoryPostRepository() *MemoryPostRepository {
	return &MemoryPostRepository{
		posts: make(map[uuid.UUID]models.Post),
	}
}

func (ur *MemoryPostRepository) Add(post models.Post) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	ur.posts[post.ID] = post
	return nil
}

func (ur *MemoryPostRepository) GetById(id uuid.UUID) (*models.Post, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	post, exists := ur.posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}

	return &post, nil
}

func (ur *MemoryPostRepository) GetAll() ([]models.Post, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	posts := make([]models.Post, 0, len(ur.posts))
	for _, post := range ur.posts {
		posts = append(posts, post)
	}

	return posts, nil
}
