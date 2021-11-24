package context

import (
	"github.com/labstack/echo/v4"

	"github.com/fr0stylo/magistras/common/pkg/services/authentication"
)

type LoggedUserContext struct {
	echo.Context
	User authentication.User
}
