package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

func TestCreateUser(t *testing.T) {
	user := models.NewUser("admin", "pass", "lewis")

	assert.Nil(t, user.Validate())
}
