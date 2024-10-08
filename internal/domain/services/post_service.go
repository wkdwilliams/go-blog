package services

import (
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/ports"
)

type IPostService interface {
	Create(title, content string, userId uuid.UUID) (*models.Post, error)
	GetById(id uuid.UUID) (*models.Post, error)
	GetAll() ([]models.Post, error)
	Delete(id uuid.UUID) error
	UpdateTitleAndContent(id uuid.UUID, title, content string) (*models.Post, error)
}

type PostService struct {
	postRepository ports.PostRepository
}

func NewPostService(r ports.PostRepository) *PostService {
	return &PostService{r}
}

func (s *PostService) Create(title, content string, userId uuid.UUID) (*models.Post, error) {
	post := models.NewPost(title, content, userId)

	if err := post.Validate(); err != nil {
		return nil, err
	}

	if err := s.postRepository.Create(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) GetById(id uuid.UUID) (*models.Post, error) {
	return s.postRepository.GetById(id)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.postRepository.GetAll()
}

func (s *PostService) Delete(id uuid.UUID) error {
	return s.postRepository.Delete(id)
}

func (s *PostService) UpdateTitleAndContent(id uuid.UUID, title, content string) (*models.Post, error) {
	post, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	post.Title = title
	post.Content = content

	if err := post.Validate(); err != nil {
		return nil, err
	}

	if err := s.postRepository.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}
