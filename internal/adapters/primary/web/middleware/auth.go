package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

func AuthenticatedUser(userService services.IUserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("goblog", c)
			if err != nil {
				return err
			}

			_userId := sess.Values["user_id"]
			if _userId == nil {
				return next(c)
			}

			userId, err := uuid.Parse(_userId.(string))
			if err != nil {
				return err
			}

			user, err := userService.GetById(userId)
			if err != nil {
				return err
			}

			c.SetRequest(c.Request().WithContext(
				context.WithValue(c.Request().Context(), "user", user),
			))

			return next(c)
		}
	}
}

func AdminAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == c.Echo().Reverse("admin-login") {
			return next(c)
		}

		sess, err := session.Get("goblog", c)
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("admin-login"))
		}

		if _, ok := sess.Values["user_id"]; !ok {
			return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("admin-login"))
		}

		return next(c)
	}

}
