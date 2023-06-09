package models

import "time"

type Post struct {
	ID          int                 `json:"id" gorm:"primary_key:auto_increment"`
	Title       string              `json:"title" form:"title" gorm:"type: varchar(255)"`
	Description string              `json:"description" form:"description" gorm:"type: varchar(255)"`
	Photos      []PhotoPostResponse `json:"photos" gorm:"foreignKey:PostID"`
	UserID      int                 `json:"user_id"`
	User        UserResponse        `json:"user"`
	CreatedAt   time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time           `json:"updated_at" gorm:"autoCreateTime"`
}

type PostResponse struct {
	ID          int                 `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Photos      []PhotoPostResponse `json:"photos" gorm:"foreignKey:PostID"`
	UserID      int                 `json:"user_id"`
	User        UserResponse        `json:"user" gorm:"foreignKey:UserID"`
}

func (PostResponse) TableName() string {
	return "posts"
}
