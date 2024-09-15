package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

type userKeyType string

const userKey userKeyType = "user" // The IDE complains if we don't do this

func AuthenticatedUser(userService services.IUserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("goblog", c)
			if err != nil {
				log.Error("failed getting session in auth middleware")
				return err
			}

			_userId := sess.Values["user_id"]
			if _userId == nil {
				log.Error("failed getting user id from session in auth middleware")
				return next(c)
			}

			userId, err := uuid.Parse(_userId.(string))
			if err != nil {
				log.Error("failed parsing the uuid in auth middleware")
				return next(c)
			}

			user, err := userService.GetById(userId)
			if err != nil {
				log.Error("failed getting the user from service in auth middleware")
				return next(c)
			}

			c.SetRequest(c.Request().WithContext(
				context.WithValue(c.Request().Context(), userKey, user),
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
