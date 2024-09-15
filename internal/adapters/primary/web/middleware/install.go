package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

func IsInstalled(userService services.IUserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			onInstallPage := c.Request().URL.Path == c.Echo().Reverse("install-page")
			onInstall := c.Request().URL.Path == c.Echo().Reverse("install")
			userCount := userService.GetTotalCount()

			if userCount == 0 {
				if !onInstallPage && !onInstall {
					c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("install-page"))
				}
			} else {
				if onInstallPage && onInstall {
					c.Redirect(http.StatusTemporaryRedirect, c.Echo().Reverse("index"))
				}
			}

			return next(c)
		}
	}
}
