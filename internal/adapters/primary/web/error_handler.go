package web

import (
	"errors"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/views"
)

func ErrorHandler(err error, c echo.Context) {
	var view templ.Component

	log.Error(err)

	if errors.Is(err, echo.ErrNotFound) {
		view = views.NotFound()
	} else if uuid.IsInvalidLengthError(err) {
		view = views.NotFound()
	} else if errors.Is(err, echo.ErrBadRequest) {
		view = views.BadRequest()
	} else {
		view = views.ServerError()
	}

	view.Render(c.Request().Context(), c.Response().Writer)
}
