package projectsdto

type ProjectRequest struct {
	Description string `json:"description"`
	Image1      string `json:"image_1" form:"image1" validate:"required"`
	Image2      string `json:"image_2" form:"image2" validate:"required"`
	Image3      string `json:"image_3" form:"image3" validate:"required"`
	Image4      string `json:"image_4" form:"image4" validate:"required"`
	Image5      string `json:"image_5" form:"image5" validate:"required"`
}