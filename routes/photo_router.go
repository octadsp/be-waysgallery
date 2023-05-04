package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func PhotoRoutes(e *echo.Group) {
	photoRepository := repositories.RepositoryPhoto(mysql.DB)

	h := handlers.HandlerPhoto(photoRepository)

	e.GET("/photos", h.FindPhotos)
	e.GET("/photo/:id", h.GetPhoto)
	e.POST("/photo", middleware.UploadImage(h.AddPhoto))
	e.PATCH("/photo/:id", middleware.UploadImage(h.UpdatePhoto))
	e.DELETE("/photo/:id", h.DeletePhoto)
}
