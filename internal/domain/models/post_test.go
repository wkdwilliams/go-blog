package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

func TestCreatePostValidation(t *testing.T) {
	userId := uuid.New()
	post := models.NewPost("first blog post", "welcome to my first post", userId)

	assert.Nil(t, post.Validate())
}
