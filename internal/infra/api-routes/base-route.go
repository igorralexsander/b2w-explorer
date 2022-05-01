package api_routes

import "github.com/labstack/echo/v4"

type BaseRoute interface {
	RegisterEndpoints(e *echo.Echo)
}
