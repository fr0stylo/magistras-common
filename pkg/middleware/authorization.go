package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/fr0stylo/magistras/common/pkg/services/authentication"
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

type Response struct {
	Error string
}

func WithAuthenticationUserContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			cookie, _ := c.Cookie("sid")
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "cookie", cookie)))

			user, err := authentication.GetAuthenticatedUser(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, &Response{Error: "Unauthorized"})
			}

			cc := authentication.LoggedUserContext{c, *user}

			return next(cc)
		}
	}
}
