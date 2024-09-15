package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web/views"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

func HandleShowInstall(c echo.Context) error {
	return views.InstallPage().Render(c.Request().Context(), c.Response().Writer)
}

type installCreateUserReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Name string `form:"name"`
}

func HandleInstall(userService services.IUserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var createUserReq installCreateUserReq

		if err := c.Bind(&createUserReq); err != nil {
			return err
		}

		if _, err := userService.CreateAccount(createUserReq.Username, createUserReq.Password, createUserReq.Name); err != nil {
			return err
		}

		return c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("admin-login"))
	}
}
