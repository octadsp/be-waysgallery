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

	e.GET("/posts", h.FindPhotos)
	e.GET("/post/:id", h.GetPhoto)
	e.POST("/post", middleware.UploadImage(h.AddPhoto))
	e.PATCH("/post/:id", middleware.UploadImage(h.UpdatePhoto))
	e.DELETE("/post/:id", h.DeletePhoto)
}
