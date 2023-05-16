package models

type User struct {
	ID       int               `json:"id" gorm:"primary_key:auto_increment"`
	Email    string            `json:"email" gorm:"type: varchar(255); unique "`
	Password string            `json:"password" gorm:"type: varchar(255)"`
	FullName string            `json:"fullName" gorm:"type: varchar(255)"`
	Greeting string            `json:"greeting" gorm:"type:varchar(255)"`
	Avatar   string            `json:"image"`
	Posts    []PostResponse    `json:"posts" gorm:"foreignKey:UserID"`
	Arts     []ArtUserResponse `json:"arts" gorm:"foreignKey:UserID"`
	Orders   []OrderUserResponse   `json:"orders" gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	ID       int               `json:"id"`
	FullName string            `json:"fullName"`
	Email    string            `json:"email"`
	Greeting string            `json:"greeting"`
	Avatar   string            `json:"image"`
	Posts    []PostResponse    `json:"posts" gorm:"foreignKey:UserID"`
	Arts     []ArtUserResponse `json:"arts" gorm:"foreignKey:UserID"`
	Orders   []OrderUserResponse   `json:"orders" gorm:"foreignKey:UserID"`
}

func (UserResponse) TableName() string {
	return "users"
}
