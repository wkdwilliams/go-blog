package web

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

type App struct {
	echo        *echo.Echo            // The router
	port        int                   // The port to listen on
	startTime   time.Time             // The time of which the server started
	userService services.IUserService // User service
	postService services.IPostService // Post service
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
	log.Printf("Shutting down after %v of up time", a.Uptime())
	return a.echo.Shutdown(ctx)
}

func (a App) Uptime() time.Duration {
	return time.Since(a.startTime)
}
