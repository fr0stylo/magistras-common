package authentication

import (
	"github.com/labstack/echo/v4"
)

type LoggedUserContext struct {
	echo.Context
	User User
}
