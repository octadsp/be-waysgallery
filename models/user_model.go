package models

import "time"

type User struct {
	ID       int               `json:"id" gorm:"primary_key:auto_increment"`
	Email    string            `json:"email" gorm:"type: varchar(255); unique "`
	Password string            `json:"password" gorm:"type: varchar(255)"`
	FullName string            `json:"fullName" gorm:"type: varchar(255)"`
	Avatar   string            `json:"image"`
	Posts    []PostResponse    `json:"posts" gorm:"foreignKey:UserID"`
	Arts     []ArtUserResponse `json:"arts" gorm:"foreignKey:UserID"`
	Orders   []OrderResponse   `json:"orders" gorm:"foreignKey:UserID"`

	//association_jointable_foreignKey untuk menyesuaikan foreignKey asing pada table sambungan Many-to-Many antara 2 struct yang berbeda
	Followers []*User   `gorm:"many2many:user_followers;association_jointable_foreignkey:follower_id" json:"-"`
	Following []*User   `gorm:"many2many:user_followers;association_jointable_foreignkey:following_id" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Avatar   string `json:"image"`
}

func (UserResponse) TableName() string {
	return "users"
}
