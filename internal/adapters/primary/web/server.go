package web

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

type App struct {
	echo        *echo.Echo
	port        int
	startTime   time.Time
	userService services.IUserService
	postService services.IPostService
}

func NewApp(userService services.IUserService, postService services.IPostService, opts ...AppOption) *App {
	s := &App{
		echo:        echo.New(),
		port:        8000,
		userService: userService,
		postService: postService,
	}

	s.echo.HTTPErrorHandler = ErrorHandler

	for _, applyOption := range opts {
		applyOption(s)
	}

	s.initAppRoutes()

	return s
}

func (a *App) Start() error {
	a.startTime = time.Now()
	return a.echo.Start(fmt.Sprintf(":%d", a.port))
}

func (a App) Stop(ctx context.Context) error {
	return a.echo.Shutdown(ctx)
}

func (a App) Uptime() time.Duration {
	return time.Since(a.startTime)
}
