package models

import "time"

type Order struct {
	ID          int          `json:"id" gorm:"primary_key:auto_increment"`
	Title       string       `json:"title" gorm:"type:varchar(255)"`
	Description string       `json:"description" gorm:"type:varchar(255)"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Price       int          `json:"price"`
	Status      string       `json:"status"`
	UserID      int          `json:"orderBy"`
	User        UserResponse `json:"user" gorm:"foreignKey:UserID"`
	OrderToID   int          `json:"orderTo"`
	OrderTo     UserResponse `json:"order_to" gorm:"foreignKey:OrderToID"`
}

type OrderResponse struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Price       int          `json:"price"`
	UserID      int          `json:"user_id"`
	User        UserResponse `json:"user" gorm:"foreignKey:UserID"`
	OrderToID   int          `json:"order_to_id"`
	OrderTo     UserResponse `json:"order_to" gorm:"foreignKey:OrderToID"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
