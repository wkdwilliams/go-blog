package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

func CreateAccount(userService services.IUserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input struct {
			Username string `json:"username"`
		}

		if err := c.Bind(&input); err != nil {
			return c.String(400, "")
		}

		resp, err := userService.CreateAccount(services.CreateAccountReq(input))
		if err != nil {
			return c.String(500, "")
		}

		var out struct {
			UserID string `json:"userId"`
		}
		out.UserID = resp.ID.String()

		return c.JSON(201, out)
	}
}
