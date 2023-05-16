package orderdto

import (
	"time"
)

type OrderResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Price       int       `json:"price"`
	VendorID    int       `json:"vendor_id"`
	ClientID    int       `json:"client_id"`
	Status      string    `json:"status"`
}
