package orderdto

import (
	"time"
)

type OrderResponse struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Price       int          `json:"price"`
	UserID      int          `json:"user_id"`
	OrderToID   int          `json:"order_to_id"`
}
