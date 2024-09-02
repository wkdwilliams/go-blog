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
	} else if errors.Is(err, echo.ErrBadRequest) {
		view = views.BadRequest()
	} else {
		view = views.ServerError()
	}

	view.Render(c.Request().Context(), c.Response().Writer)
}
