package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

func TestCreatePost(t *testing.T) {
	userId := uuid.New()
	post := models.NewPost("first blog post", "welcome to my first post", userId)

	if uuid, err := uuid.Parse(post.ID.String()); err != nil {
		t.Fatalf("uuid: %s is not valid", uuid.String())
	}

	if post.Title != "first blog post" {
		t.Fatalf("title: %s is not valid", post.Title)
	}

	if post.Content != "welcome to my first post" {
		t.Fatalf("content: %s is not valid", post.Content)
	}
}
