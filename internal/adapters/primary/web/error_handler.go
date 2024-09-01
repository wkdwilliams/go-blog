package web

import (
	"errors"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/views"
)

func ErrorHandler(err error, c echo.Context) {
	var view templ.Component

	if errors.Is(err, echo.ErrNotFound) {
		view = views.NotFound()
	}
	if errors.Is(err, echo.ErrInternalServerError) {
		view = views.ServerError()
	}
	if errors.Is(err, echo.ErrBadRequest) {
		view = views.BadRequest()
	}

	views.Main(view).Render(c.Request().Context(), c.Response().Writer)
}
