package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/admin/login" {
			return next(c)
		}

		sess, err := session.Get("goblog", c)
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
		}

		if _, ok := sess.Values["user_id"]; !ok {
			return c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
		}

		return next(c)
	}

}
