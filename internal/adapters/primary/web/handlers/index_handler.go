package handlers

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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

		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}

		fmt.Println(sess.Values["foo"])
		a := views.Index(posts)
		return a.Render(c.Request().Context(), c.Response().Writer)
	}
}

func CreateCookie(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values["foo"] = "bar"

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.NoContent(200)
}
