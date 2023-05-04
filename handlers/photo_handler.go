package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	photodto "waysgallery/dto/photo"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	repository "waysgallery/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerPhoto struct {
	PhotoRepository repository.PhotoRepository
}

func HandlerPhoto(PhotoRepository repository.PhotoRepository) *handlerPhoto {
	return &handlerPhoto{PhotoRepository}
}

func (h *handlerPhoto) FindPhotos(c echo.Context) error {
	photos, err := h.PhotoRepository.FindPhotos()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: photos})
}

func (h *handlerPhoto) GetPhoto(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	photo, err := h.PhotoRepository.GetPhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: photo})
}

func (h *handlerPhoto) AddPhoto(c echo.Context) error {
	imageFile := c.Get("imageFile").(string)
	post_id, _ := strconv.Atoi(c.FormValue("post_id"))

	request := photodto.PhotoRequest{
		Image:  imageFile,
		PostID: post_id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "waysgallery"})

	if err != nil {
		fmt.Println(err.Error())
	}

	photo := models.Photo{
		Image:  resp.SecureURL,
		PostID: request.PostID,
	}

	photo, err = h.PhotoRepository.AddPhoto(photo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	photo, _ = h.PhotoRepository.GetPhoto(photo.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: photo})
}

func (h *handlerPhoto) UpdatePhoto(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	//GET IMAGE FILE
	imageFile := c.Get("imageFile").(string)
	post_id, _ := strconv.Atoi(c.FormValue("post_id"))

	request := photodto.PhotoRequest{
		Image:  imageFile,
		PostID: post_id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{Folder: "waysgallery"})

	if err != nil {
		fmt.Println(err.Error())
	}

	photo, err := h.PhotoRepository.GetPhoto(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Image != "" {
		photo.Image = resp.SecureURL
	}

	if request.PostID != 0 {
		photo.PostID = request.PostID
	}

	updatedPhoto, err := h.PhotoRepository.UpdatePhoto(photo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: updatedPhoto})
}

func (h *handlerPhoto) DeletePhoto(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	photo, err := h.PhotoRepository.GetPhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.PhotoRepository.DeletePhoto(photo, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
