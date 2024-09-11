package web

import (
	"embed"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/handlers"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/middleware"
)

//go:embed static
var static embed.FS

func (a *App) initAppRoutes() {
	// We need this to make static files (js, css) public
	a.echo.StaticFS("/static", echo.MustSubFS(static, "static"))
	a.echo.Pre(echoMiddleware.RemoveTrailingSlash())

	main := a.echo.Group("")
	{
		// Middleware for everything
		main.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SECRET")))))
		main.Use(middleware.AuthenticatedUser(a.userService))

		// Health check
		main.GET("/health", func(c echo.Context) error {
			return c.String(200, "ok")
		}).Name = "health"

		// Handle the index request e.g. http://host/
		main.GET("", handlers.IndexHandler(a.postService)).Name = "index"

		// Group for the admin routes e.g. http://host/admin/*
		admin := main.Group("/admin", middleware.AdminAuthorized)
		{
			// Handle the admin index request e.g. http://host/admin
			admin.GET("", handlers.AdminIndexHandler).Name = "admin-index"

			// Handle the post create request
			admin.POST("/post", handlers.AdminPostCreateHandler(a.postService)).Name = "admin-post"

			// Handle the post delete request
			admin.GET("/post/delete/:id", handlers.AdminPostDeleteHandler(a.postService)).Name = "admin-post-delete"

			// Handle the login request to show login form
			admin.GET("/login", handlers.AdminLoginHandler).Name = "admin-login"

			// Handle the login request to authenticate the user
			admin.POST("/login", handlers.AdminTryLoginHandler(a.userService)).Name = "admin-login-try"

			// Handle the logout request
			admin.GET("/logout", handlers.AdminLogout).Name = "admin-logout"

			// Handle the user account create request
			admin.POST("/users", handlers.CreateAccount(a.userService)).Name = "admin-user-create"
		}
	}
}
