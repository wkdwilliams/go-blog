package web

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

type App struct {
	echo        *echo.Echo
	UserService services.IUserService
	PostService services.IPostService
	port        int
}

func NewApp(userService services.IUserService, postService services.IPostService, opts ...AppOption) *App {
	s := &App{
		echo:        echo.New(),
		UserService: userService,
		PostService: postService,
		port:        8000,
	}

	s.echo.HTTPErrorHandler = ErrorHandler

	for _, applyOption := range opts {
		applyOption(s)
	}

	s.initAppRoutes()

	return s
}

func (a App) GetAuthenticatedUser(c echo.Context) (*models.User, error) {
	sess, err := session.Get("goblog", c)
	if err != nil {
		return nil, err
	}
	if _, ok := sess.Values["user_id"]; !ok {
		return nil, errors.New("not authenticated")
	}

	return a.UserService.GetById(sess.Values["user_id"].(uuid.UUID))
}

func (a App) Start() error {
	return a.echo.Start(fmt.Sprintf(":%d", a.port))
}

func (a App) Stop(ctx context.Context) error {
	return a.echo.Shutdown(ctx)
}
