package component

import (
	"github.com/labstack/echo/v4"
)

type Component interface {
	Name() string
	RegisterEventHandler()
	RegisterRouter(e *echo.Echo, m ...echo.MiddlewareFunc)
}
