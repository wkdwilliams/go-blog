package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/views"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

func IndexHandler(postService services.IPostService) echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := postService.GetAll()

		if err != nil {
			return err
		}

		return views.Main(views.Home(posts)).Render(c.Request().Context(), c.Response().Writer)
	}
}
