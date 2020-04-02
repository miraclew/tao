package component

import (
	"github.com/labstack/echo/v4"
)

// A basic locatable component
type Component interface {
	Name() string
}

// A http service component
type Service interface {
	Component
	RegisterEventHandler()
	RegisterRouter(e *echo.Echo, m ...echo.MiddlewareFunc)
}
