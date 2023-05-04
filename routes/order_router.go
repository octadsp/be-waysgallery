package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Group) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)

	h := handlers.HandlerOrder(orderRepository)

	e.GET("/orders", h.FindOrders)
	e.GET("/order/:id", h.GetOrder)
	e.POST("/order", h.AddOrder)
	e.PATCH("/order/:id", h.UpdateOrder)
	e.DELETE("/order/:id", h.DeleteOrder)
}
