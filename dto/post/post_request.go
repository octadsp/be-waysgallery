package postdto

type PostRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	UserID      int    `json:"user_id" form:"user_id"`
}
