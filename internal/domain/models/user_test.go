package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

func TestCreateUser(t *testing.T) {
	user := models.NewUser("lewis")

	if uuid, err := uuid.Parse(user.ID.String()); err != nil {
		t.Fatalf("uuid: %s is not valid", uuid.String())
	}

	if user.Username != "lewis" {
		t.Fatalf("username: %s is not valid", user.Username)
	}
}
