package models

import "time"

type Post struct {
	ID          int                 `json:"id" gorm:"primary_key:auto_increment"`
	Title       string              `json:"title" form:"title" gorm:"type: varchar(255)"`
	Description string              `json:"description" form:"description" gorm:"type: varchar(255)"`
	Photos      []PhotoPostResponse `json:"photo" gorm:"foreignKey:PostID"`
	UserID      int                 `json:"user_id"`
	User        UserResponse        `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time           `json:"updated_at" gorm:"autoCreateTime"`
}

type PostResponse struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Photo       string       `json:"photo"`
	UserID      int          `json:"user_id"`
	User        UserResponse `json:"user"`
}

func (PostResponse) TableName() string {
	return "posts"
}
