package context_helper

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
)

func GetUserFromEchoContext(c echo.Context) *models.User {
	return c.Request().Context().Value("user").(*models.User)
}

func GetUserFromContext(c context.Context) *models.User {
	if user := c.Value("user"); user != nil {
		return user.(*models.User)
	}

	return nil
}

func UserIsLoggedInFromContext(ctx context.Context) bool {
	return GetUserFromContext(ctx) != nil
}
