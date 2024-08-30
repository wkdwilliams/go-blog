package web

import (
	"embed"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/handlers"
)

//go:embed static
var static embed.FS

func (a *App) initAppRoutes() {
	a.echo.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	a.echo.GET("/", handlers.IndexHandler(a.PostService))
	a.echo.GET("/cookie", handlers.CreateCookie)
	a.echo.StaticFS("/static", echo.MustSubFS(static, "static"))
	a.echo.POST("/users", handlers.CreateAccount(a.UserService))
}
