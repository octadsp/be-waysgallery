package orderdto

import "time"

type OrderRequest struct {
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	StartDate   time.Time `json:"start_date" form:"start_date" `
	EndDate     time.Time `json:"end_date" form:"end_date"`
	Price       int       `json:"price" form:"price"`
	UserID      int       `json:"orderBy"`
	OrderToID   int       `json:"orderTo"`
}
