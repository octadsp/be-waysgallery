package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func ProjectRoutes(e *echo.Group) {
	projectRepository := repositories.RepositoryProject(mysql.DB)
	h := handlers.HandlerProject(projectRepository)

	e.GET("/projects", h.FindProjects)
	e.GET("/project/:id", h.GetProject)
	e.POST("/project/:order_id", middleware.Auth(middleware.UploadImage(h.CreateProject)))
}