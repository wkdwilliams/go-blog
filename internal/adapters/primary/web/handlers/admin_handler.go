package handlers

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/views"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/pkg/context_helper"
)

func AdminIndexHandler(c echo.Context) error {
	return views.Admin(false, nil).Render(c.Request().Context(), c.Response().Writer)
}

func AdminLoginHandler(c echo.Context) error {
	return views.AdminLogin(false).Render(c.Request().Context(), c.Response().Writer)
}

// The request for creating a new blog post
type createPostRequest struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

func (a createPostRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.Content, validation.Required, validation.Length(1, 10000)),
	)
}

func AdminPostCreateHandler(postService services.IPostService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload createPostRequest
		if err := c.Bind(&payload); err != nil {
			return err
		}

		if err := payload.Validate(); err != nil {
			return views.Admin(false, err).Render(c.Request().Context(), c.Response().Writer)
		}

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
			fmt.Println("deleting failed")
			return err
		}

		return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("index"))
	}
}

func AdminPostEditHandler(postService services.IPostService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			return err
		}

		post, err := postService.GetById(id)

		if err != nil {
			return err
		}

		return views.AdminPostEdit(false, nil, post).Render(c.Request().Context(), c.Response().Writer)
	}
}

func AdminPostTryEditHandler(postService services.IPostService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			return err
		}

		var editPostRequest createPostRequest

		if err := c.Bind(&editPostRequest); err != nil {
			return err
		}

		post, err := postService.UpdateTitleAndContent(id, editPostRequest.Title, editPostRequest.Content)

		if err != nil {
			return err
		}

		return views.AdminPostEdit(true, nil, post).Render(c.Request().Context(), c.Response().Writer)
	}
}

type loginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (a loginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.Password, validation.Required, validation.Length(1, 0)),
	)
}

func AdminTryLoginHandler(userService services.IUserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		loginRequest := &loginRequest{}

		if err := c.Bind(loginRequest); err != nil {
			return err
		}

		if err := loginRequest.Validate(); err != nil {
			return views.AdminLogin(true).Render(c.Request().Context(), c.Response().Writer)
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
