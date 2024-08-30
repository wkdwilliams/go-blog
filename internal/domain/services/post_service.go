package services

import (
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
)

type IPostService interface {
	Create(title, content string, userId uuid.UUID) (*models.Post, error)
	GetById(id uuid.UUID) (*models.Post, error)
	GetAll() (*[]models.Post, error)
}

type PostService struct {
	postRepository ports.PostRepository
}

func NewPostService(r ports.PostRepository) *PostService {
	return &PostService{r}
}

func (s *PostService) Create(title, content string, userId uuid.UUID) (*models.Post, error) {
	post := models.NewPost(title, content, userId)

	if err := s.postRepository.Add(post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) GetById(id uuid.UUID) (*models.Post, error) {
	return s.postRepository.GetById(id)
}

func (s *PostService) GetAll() (*[]models.Post, error) {
	return s.postRepository.GetAll()
}
