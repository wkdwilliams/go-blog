package web

import (
	"embed"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/handlers"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/middleware"
)

//go:embed static
var static embed.FS

func (a *App) initAppRoutes() {
	a.echo.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SECRET")))))

	a.echo.StaticFS("/static", echo.MustSubFS(static, "static"))

	admin := a.echo.Group("/admin", middleware.IsLoggedIn())
	admin.GET("", handlers.AdminIndexHandler())
	admin.POST("/post", handlers.AdminPostCreateHandler(a.PostService))
	admin.GET("/login", handlers.AdminLoginHandler())
	admin.POST("/login", handlers.AdminTryLoginHandler(a.UserService))

	a.echo.GET("/", handlers.IndexHandler(a.PostService))
	a.echo.POST("/users", handlers.CreateAccount(a.UserService))
}
