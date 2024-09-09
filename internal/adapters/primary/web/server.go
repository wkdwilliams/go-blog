package web

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

type App struct {
	Echo        *echo.Echo
	port        int
	startTime   time.Time
	UserService services.IUserService
	PostService services.IPostService
}

func NewApp(userService services.IUserService, postService services.IPostService, opts ...AppOption) *App {
	s := &App{
		Echo:        echo.New(),
		port:        8000,
		UserService: userService,
		PostService: postService,
	}

	s.Echo.HTTPErrorHandler = ErrorHandler

	for _, applyOption := range opts {
		applyOption(s)
	}

	s.initAppRoutes()

	return s
}

func (a *App) Start() error {
	a.startTime = time.Now()
	return a.Echo.Start(fmt.Sprintf(":%d", a.port))
}

func (a App) Stop(ctx context.Context) error {
	return a.Echo.Shutdown(ctx)
}

func (a App) Uptime() time.Duration {
	return time.Since(a.startTime)
}
