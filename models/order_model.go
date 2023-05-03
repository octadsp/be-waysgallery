package models

import "time"

type Order struct {
	ID          int          `json:"id" gorm:"primary_key:auto_increment"`
	Title       string       `json:"title" gorm:"type:varchar(255)"`
	Description string       `json:"description" gorm:"type:varchar(255)"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Price       int          `json:"price"`
	OrderByID   int          `json:"order_by_id"`
	OrderBy     UserResponse `json:"order_by"`
	OrderToID   int          `json:"order_to_id"`
	OrderTo     UserResponse `json:"order_to"`
}

type OrderResponse struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Price       int          `json:"price"`
	OrderBy     UserResponse `json:"order_by"`
	OrderTo     UserResponse `json:"order_to"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
