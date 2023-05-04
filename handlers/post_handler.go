package handlers

import (
	"net/http"
	"strconv"
	postdto "waysgallery/dto/post"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	repository "waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerPost struct {
	PostRepository repository.PostRepository
}

func HandlerPost(PostRepository repository.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

func (h *handlerPost) FindPosts(c echo.Context) error {
	posts, err := h.PostRepository.FindPosts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: posts})
}

func (h *handlerPost) GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: post})
}

func (h *handlerPost) AddPost(c echo.Context) error {
	request := new(postdto.PostRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user_id, _ := strconv.Atoi(c.FormValue("user_id"))

	post := models.Post{
		Title:       request.Title,
		Description: request.Description,
		UserID:      user_id,
	}

	data, err := h.PostRepository.AddPost(post)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerPost) UpdatePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(postdto.PostRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	user, err := h.PostRepository.GetPost(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		user.Title = request.Title
	}

	if request.Description != "" {
		user.Description = request.Description
	}

	if request.UserID != 0 {
		user.UserID = request.UserID
	}

	data, err := h.PostRepository.UpdatePost(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerPost) DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.PostRepository.DeletePost(post, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
