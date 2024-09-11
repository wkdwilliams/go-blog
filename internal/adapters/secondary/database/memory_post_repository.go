package database

import (
	"sort"
	"sync"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
)

type MemoryPostRepository struct {
	posts map[uuid.UUID]*models.Post
	mu    sync.RWMutex
}

func NewMemoryPostRepository() *MemoryPostRepository {
	return &MemoryPostRepository{
		posts: make(map[uuid.UUID]*models.Post),
	}
}

func (pr *MemoryPostRepository) Add(p models.Post) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	// If the post ID is not set, generate a new one
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	pr.posts[p.ID] = &p
	return nil
}

func (pr *MemoryPostRepository) GetById(id uuid.UUID) (*models.Post, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	post, exists := pr.posts[id]
	if !exists {
		return nil, ports.ErrRecordNotFound
	}
	return post, nil
}

func (pr *MemoryPostRepository) GetAll() ([]models.Post, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	posts := make([]models.Post, 0, len(pr.posts))
	for _, post := range pr.posts {
		posts = append(posts, *post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})

	return posts, nil
}

func (pr *MemoryPostRepository) Delete(id uuid.UUID) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	if _, exists := pr.posts[id]; !exists {
		return ports.ErrRecordNotFound
	}
	delete(pr.posts, id)
	return nil
}
