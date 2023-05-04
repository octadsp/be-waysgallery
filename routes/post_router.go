package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Group) {
	PostRepository := repositories.RepositoryPost(mysql.DB)

	h := handlers.HandlerPost(PostRepository)

	e.GET("/posts", h.FindPosts)
	e.GET("/post/:id", h.GetPost)
	e.POST("/post", h.AddPost)
	e.PATCH("/post/:id", middleware.UploadImage(h.UpdatePost))
	e.DELETE("/post/:id", middleware.UploadImage(h.DeletePost))
}
