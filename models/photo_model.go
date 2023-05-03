package models

import "time"

type Photo struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	Image     string       `json:"image" gorm:"type:varchar(255)"`
	PostID    int          `json:"post_id"`
	Post      PostResponse `json:"post"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoCreateTime"`
}

type PhotoPostResponse struct {
	ID     int          `json:"id"`
	Image  string       `json:"image"`
	PostID int          `json:"post_id"`
	Post   PostResponse `json:"post"`
}

func (PhotoPostResponse) TableName() string {
	return "photos"
}
