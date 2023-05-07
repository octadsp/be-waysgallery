package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func ArtRoutes(e *echo.Group) {
	artRepository := repositories.RepositoryArt(mysql.DB)

	h := handlers.HandlerArt(artRepository)

	e.GET("/arts", h.FindArts)
	e.GET("/art/:id", h.GetArt)
	e.POST("/art", middleware.UploadImage(h.AddArt))
	e.PATCH("/art/:id", middleware.UploadImage(h.UpdateArt))
	e.DELETE("/art/:id", h.DeleteArt)
}
