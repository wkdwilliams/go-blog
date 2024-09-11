package web

import (
	"errors"

	"github.com/a-h/templ"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/views"
)

type ApiError struct {
	Error error `json:"error"`
}

func ErrorHandler(err error, c echo.Context) {
	var view templ.Component

	log.Error(err)

	if errors.Is(err, echo.ErrNotFound) {
		view = views.NotFound()
	} else if uuid.IsInvalidLengthError(err) {
		view = views.NotFound()
	} else if _, ok := err.(validation.Errors); ok {
		view = views.BadRequest()
		//c.JSON(http.StatusBadRequest, ApiError{err}) // <- for API
		//return
	} else if errors.Is(err, echo.ErrBadRequest) {
		view = views.BadRequest()
	} else {
		view = views.ServerError()
	}

	view.Render(c.Request().Context(), c.Response().Writer)
}
