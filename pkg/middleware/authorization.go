package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

func Authorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, _ := c.Cookie("sid")
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "cookie", cookie)))
			return next(c)
		}
	}
}
