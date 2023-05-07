package artsdto

type ArtRequest struct {
	Image  string `json:"image" form:"image" validate:"required"`
	UserID int    `json:"user_id" form:"post_id"`
}