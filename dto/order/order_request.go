package orderdto

import "time"

type OrderRequest struct {
	Title       string    `json:"title" form:"title" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	StartDate   time.Time `json:"start_date" form:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" form:"end_date" validate:"required"`
	Price       int       `json:"price" form:"price" validate:"required"`
	UserID      int       `json:"user_id" validate:"required"`
	OrderToID   int       `json:"order_to_id" validate:"required"`
}
