package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Group) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerOrder(orderRepository, userRepository)

	e.GET("/orders", middleware.Auth(h.FindOrders))
	e.GET("/order/:id", middleware.Auth(h.GetOrder))
	e.POST("/order/:vendor_id", middleware.Auth(h.CreateOrder))
	e.PATCH("/order/:id", middleware.Auth(h.UpdateOrderStatus))
	e.POST("/notification", h.Notification)
}
