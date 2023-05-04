package photodto

type PhotoRequest struct {
	Image  string `json:"image" form:"image" validate:"required"`
	PostID int    `json:"post_id" form:"post_id" validate:"required"`
}