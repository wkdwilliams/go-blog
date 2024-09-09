package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/views"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/pkg/context_helper"
	"github.com/wkdwilliams/go-blog/pkg/validator"
)

func AdminIndexHandler(c echo.Context) error {
	return views.Admin(false, nil).Render(c.Request().Context(), c.Response().Writer)
}

func AdminLoginHandler(c echo.Context) error {
	return views.AdminLogin(false).Render(c.Request().Context(), c.Response().Writer)
}

type createPostRequest struct {
	Title   string `form:"title" validate:"required,max=1"`
	Content string `form:"content" validate:"required"`
}

func AdminPostCreateHandler(postService services.IPostService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload createPostRequest
		if err := c.Bind(&payload); err != nil {
			return err
		}

		if err := validator.Validate(&payload); err != nil {
			return err
		}

		// if err := c.Validate(payload); err != nil {
		// 	return views.Admin(
		// 		false,
		// 		validator.ParseErrors(&payload, err, "form"),
		// 	).Render(c.Request().Context(), c.Response().Writer)
		// }

		if _, err := postService.Create(
			payload.Title,
			payload.Content,
			context_helper.GetUserFromEchoContext(c).ID,
		); err != nil {
			return err
		}

		return views.Admin(true, nil).Render(c.Request().Context(), c.Response().Writer)
	}
}

func AdminPostDeleteHandler(postService services.IPostService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			return err
		}

		if err := postService.Delete(id); err != nil {
			return err
		}

		return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("index"))
	}
}

type loginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func AdminTryLoginHandler(userService services.IUserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		loginRequest := &loginRequest{}

		if err := c.Bind(loginRequest); err != nil {
			return err
		}

		user, err := userService.Authenticate(loginRequest.Username, loginRequest.Password)

		if err != nil {
			return views.AdminLogin(true).Render(c.Request().Context(), c.Response().Writer)
		}

		sess, err := session.Get("goblog", c)

		if err != nil {
			return err
		}

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		sess.Values["user_id"] = user.ID.String()

		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin-index"))
	}
}

func AdminLogout(c echo.Context) error {
	sess, err := session.Get("goblog", c)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	err = sess.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("index"))
}
