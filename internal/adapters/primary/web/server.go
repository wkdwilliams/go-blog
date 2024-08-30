package web

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
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

	for _, applyOption := range opts {
		applyOption(s)
	}

	s.initAppRoutes()

	return s
}

func (a App) Start() error {
	return a.echo.Start(fmt.Sprintf(":%d", a.port))
}

func (a App) Stop(ctx context.Context) error {
	return a.echo.Shutdown(ctx)
}
