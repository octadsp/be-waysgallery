package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	artsdto "waysgallery/dto/art"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	repository "waysgallery/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerArt struct {
	ArtRepository repository.ArtRepository
}

func HandlerArt(ArtRepository repository.ArtRepository) *handlerArt {
	return &handlerArt{ArtRepository}
}

func (h *handlerArt) FindArts(c echo.Context) error {
	arts, err := h.ArtRepository.FindArts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: arts})
}

func (h *handlerArt) GetArt(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	art, err := h.ArtRepository.GetArt(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: art})
}

func (h *handlerArt) AddArt(c echo.Context) error {
	imageFile := c.Get("imageFile").(string)
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := artsdto.ArtRequest{
		Image:  imageFile,
		UserID: user_id,
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

	art := models.Art{
		Image:  resp.SecureURL,
		UserID: request.UserID,
	}

	art, err = h.ArtRepository.AddArt(art)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	art, _ = h.ArtRepository.GetArt(art.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: art})
}

func (h *handlerArt) UpdateArt(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	//GET IMAGE FILE
	imageFile := c.Get("imageFile").(string)
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))

	request := artsdto.ArtRequest{
		Image:  imageFile,
		UserID: user_id,
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

	art, err := h.ArtRepository.GetArt(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Image != "" {
		art.Image = resp.SecureURL
	}

	if request.UserID != 0 {
		art.UserID = request.UserID
	}

	updateArt, err := h.ArtRepository.UpdateArt(art)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: updateArt})
}

func (h *handlerArt) DeleteArt(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	art, err := h.ArtRepository.GetArt(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ArtRepository.DeleteArt(art, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
