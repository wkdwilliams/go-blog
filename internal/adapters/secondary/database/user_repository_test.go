package database_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/database"
	"github.com/wkdwilliams/go-blog/internal/infrastructure"
	"github.com/wkdwilliams/go-blog/pkg/hashing"
)

func TestGetAllUsers(t *testing.T) {
	db, mock := infrastructure.NewMysqlMock()

	userRepo := database.NewUserRepository(db)

	hashed, _ := hashing.HashPassword("pass")

	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).
		AddRow(uuid.New(), "new", hashed, time.Now(), time.Now())

	mock.ExpectQuery("SELECT (.*)").WillReturnRows(rows)

	allUsers, err := userRepo.GetAll()

	if err != nil {
		t.Fatal(err)
	}

	if len(allUsers) != 1 {
		t.Fatal("unexpected length of users")
	}

	if allUsers[0].Username != "new" {
		t.Fatal("unexpected username")
	}

	if allUsers[0].Password == "" {
		t.Fatal("unexpected empty password")
	}

	if !hashing.VerifyPassword("pass", hashed) {
		t.Fatal("wrong password")
	}
}
