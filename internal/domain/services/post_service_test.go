package services_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/mysql"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/internal/infrastructure"
	"testing"
)

func TestGetUserById(t *testing.T) {
	db, mock := infrastructure.NewMysqlMock()

	departmentRows := sqlmock.NewRows([]string{"id", "username"}).
		AddRow(1, "Engineering").
		AddRow(2, "Marketing")

	mock.ExpectQuery("SELECT id, department_name FROM departments").
		WillReturnRows(departmentRows)

	userRepo := mysql.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
}
