package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AuthRoutes(e)
	PostRoutes(e)
	PhotoRoutes(e)
	OrderRoutes(e)
}
