package models

type Art struct {
	ID     int          `json:"id" gorm:"primary_key:auto_increment"`
	Image  string       `json:"image" form:"image"`
	UserID int          `json:"user_id"`
	User   UserResponse `json:"user" gorm:"foreignKey:UserID"`
}

type ArtUserResponse struct {
	ID     int          `json:"id"`
	Image  string       `json:"image"`
	UserID int          `json:"user_id"`
	User   UserResponse `json:"user"`
}

func (ArtUserResponse) TableName() string {
	return "orders"
}
